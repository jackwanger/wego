package hardict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistInvalidWord(t *testing.T) {
	assert.Equal(t, true, ExistInvalidWord("测试封杀"))
}

func TestReplaceInvalidWords(t *testing.T) {
	assert.Equal(t, "测试**", ReplaceInvalidWords("测试封杀"))
}
