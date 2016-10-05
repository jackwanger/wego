package main

import (
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/huichen/sego"
)

var segmenter sego.Segmenter

const DICT_PATH = "dict.txt"

func main() {
	segmenter.LoadDictionary(DICT_PATH)

	router := gin.Default()
	router.POST("/validate", validateEndPoint)
	router.POST("/filter", filterEndPoint)
	router.Run(":3001")
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
