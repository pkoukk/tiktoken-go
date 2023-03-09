package tiktoken

import (
	"regexp"
	"testing"

	"github.com/dlclark/regexp2"
	"github.com/stretchr/testify/assert"
)

func TestRegex2Func(t *testing.T) {
	ass := assert.New(t)
	pattern := `[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`
	re := regexp.MustCompile(pattern)
	re2 := regexp2.MustCompile(pattern, regexp2.None)

	words := []string{
		"this is my email hi@google.com,and this is john's email world@outlook.com",
		"hi@google.com is email for google",
		"outlook email world@outlook.com is work for microsoft",
	}

	for _, word := range words {
		ass.ElementsMatch(re.FindStringIndex(word), findRegex2StringIndex(word, re2))
		ass.ElementsMatch(re.FindAllStringSubmatchIndex(word, -1), findRegex2AllStringMatchIndex(word, re2))
		ass.Equal(re.FindString(word), findRegex2StringMatch(word, re2))
	}
}
