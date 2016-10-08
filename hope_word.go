package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/huichen/sego"
)

var port int
var segmenter sego.Segmenter

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()
}

func main() {
	files, err := prepareDict()
	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	defer removeDict(files)
	dict := strings.Join(files, ",")
	segmenter.LoadDictionary(dict)

	r := gin.Default()
	r.POST("/validate", validateEndPoint)
	r.POST("/filter", filterEndPoint)
	r.Run(fmt.Sprintf(":%d", port))
}

func validateEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	segments := segmenter.Segment([]byte(text))
	if existInvalidWord(segments) {
		c.JSON(200, gin.H{"result": "false"})
	} else {
		c.JSON(200, gin.H{"result": "true"})
	}
}

func existInvalidWord(segments []sego.Segment) bool {
	for _, seg := range segments {
		token := seg.Token()
		if token.Frequency() > 1 {
			return true
		}
	}
	return false
}

func filterEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	segments := segmenter.Segment([]byte(text))
	text = replaceInvalidWords(segments, text)
	c.JSON(200, gin.H{"result": text})
}

func replaceInvalidWords(segments []sego.Segment, text string) string {
	for _, seg := range segments {
		token := seg.Token()
		if token.Frequency() > 1 {
			oldText := token.Text()
			newText := strings.Repeat("*", utf8.RuneCountInString(oldText))
			text = regexp.
				MustCompile(fmt.Sprintf("(?i)%s", oldText)).
				ReplaceAllLiteralString(text, newText)
		}
	}
	return text
}

//go:generate go-bindata -prefix "dict/" -pkg main -o dict.go dict/
func prepareDict() ([]string, error) {
	var files = make([]string, 0)
	for _, v := range AssetNames() {
		data, err := Asset(v)
		if err != nil {
			return nil, err
		}
		f, err := saveTempFile(v, data)
		if err != nil {
			return nil, err
		}
		files = append(files, f.Name())
	}
	return files, nil
}

func saveTempFile(asset string, data []byte) (*os.File, error) {
	tmpfile, err := ioutil.TempFile("", asset)
	if err != nil {
		return nil, err
	}
	if _, err := tmpfile.Write(data); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}

func removeDict(files []string) {
	for _, f := range files {
		os.Remove(f)
	}
}
