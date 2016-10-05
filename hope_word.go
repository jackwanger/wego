package main

import (
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/huichen/sego"
)

const DICT_PATH = "dict.txt"

func main() {
	r := gin.Default()
	r.Use(Segmenter())
	r.POST("/validate", validateEndPoint)
	r.POST("/filter", filterEndPoint)
	r.Run(":8000")
}

func Segmenter() gin.HandlerFunc {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(DICT_PATH)
	return func(c *gin.Context) {
		c.Set("Segmenter", segmenter)
		c.Next()
	}
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
	text = ReplaceInvalidWords(segments, text)
	c.JSON(200, gin.H{"result": text})
}
func ReplaceInvalidWords(segments []sego.Segment, text string) string {
	for _, seg := range segments {
		token := seg.Token()
		if token.Frequency() > 1 {
			oldText := token.Text()
			newText := strings.Repeat("*", utf8.RuneCountInString(oldText))
			text = strings.Replace(text, oldText, newText, -1)
		}
	}
	return text
}
