//go:generate go-bindata -prefix "assets/" -pkg hardict -o assets.go assets/

package hardict

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/huichen/sego"
)

var segmenter sego.Segmenter

func init() {
	fmt.Println("Loading dict...")
	load(&segmenter)
}

func getSegments(text string) []sego.Segment {
	return segmenter.Segment([]byte(text))
}

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
