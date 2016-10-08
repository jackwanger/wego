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
)

var port int
var dict string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&dict, "dict", "dict.txt", "dict files, seperated by comma")
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()
}

func main() {
	r := gin.Default()

	// Middleware
	r.Use(func(c *gin.Context) {
		var segmenter sego.Segmenter
		segmenter.LoadDictionary(dict)
		c.Set("Segmenter", segmenter)
		c.Next()
	})

	// API
	r.POST("/validate", validateEndPoint)
	r.POST("/filter", filterEndPoint)
	r.Run(fmt.Sprintf(":%d", port))
}

func validateEndPoint(c *gin.Context) {
	text := c.PostForm("message")
	s := c.MustGet("Segmenter").(sego.Segmenter)
	segments := s.Segment([]byte(text))
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
	s := c.MustGet("Segmenter").(sego.Segmenter)
	segments := s.Segment([]byte(text))
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
