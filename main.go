package main

import (
	"flag"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/huichen/sego"
	"github.com/repong/hope_word/dict"
)

var port int
var segmenter sego.Segmenter

func main() {
	version()

	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
	dict.Load(&segmenter)

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
