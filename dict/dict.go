package dict

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/repong/sego"
)

var segmenter sego.Segmenter

// Load dictionaries from dictPath
func Load(dictPath string) {
	segmenter.LoadDictionary(dictPath)
}

// ExistInvalidWord Check if text contains words defined in dictionary
func ExistInvalidWord(text string) bool {
	segments := getSegments(text)
	for _, seg := range segments {
		token := seg.Token()
		if token.Frequency() > 1 {
			return true
		}
	}
	return false
}

// ReplaceInvalidWords Replace words defineds in dictionary
func ReplaceInvalidWords(text string) string {
	segments := getSegments(text)
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

func getSegments(text string) []sego.Segment {
	return segmenter.Segment([]byte(text))
}
