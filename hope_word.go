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

var segmenter sego.Segmenter
var port int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()

	if err := prepareDict(); err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	dict := strings.Join(AssetNames(), ",")
	segmenter.LoadDictionary(dict)
}

//go:generate go-bindata -prefix "dict/" -pkg main -o dict.go dict/
func prepareDict() error {
	for _, v := range AssetNames() {
		data, _ := Asset(v)
		err := ioutil.WriteFile(v, []byte(data), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	r := gin.Default()
	r.POST("/validate", validateEndPoint)
	r.POST("/filter", filterEndPoint)
	r.Run(fmt.Sprintf(":%d", port))
}

func validateEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	segments := segmenter.Segment([]byte(text))
	if IsContainInvalidWord(segments) {
		c.JSON(200, gin.H{"result": "false"})
	} else {
		c.JSON(200, gin.H{"result": "true"})
	}
}

func IsContainInvalidWord(segments []sego.Segment) bool {
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
