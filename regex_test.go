package tiktoken

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex2Func(t *testing.T) {
	ass := assert.New(t)
	pattern := `[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`
	re := regexp.MustCompile(pattern)
	re2 := regexp.MustCompile(pattern)

	words := []string{
		"this is my email hi@google.com,and this is john's email world@outlook.com",
		"hi@google.com is email for google",
		"outlook email world@outlook.com is work for microsoft",
	}

	for _, word := range words {
		// Compare byte indices converted to rune indices
		byteIndices := re.FindStringIndex(word)
		if byteIndices != nil {
			byteToRune := byteToRuneMap(word)
			runeStart, runeEnd := byteIndicesToRuneIndicesWithMap(byteToRune, byteIndices[0], byteIndices[1])
			expectedRuneIndices := []int{runeStart, runeEnd}
			ass.ElementsMatch(expectedRuneIndices, findRegex2StringIndex(word, re2))
		} else {
			ass.Nil(findRegex2StringIndex(word, re2))
		}

		// For FindAllStringIndex, we need to convert all byte indices to rune indices
		allByteIndices := re.FindAllStringIndex(word, -1)
		if allByteIndices != nil {
			byteToRune := byteToRuneMap(word)
			expectedRuneIndices := make([][]int, len(allByteIndices))
			for i, byteIndices := range allByteIndices {
				runeStart, runeEnd := byteIndicesToRuneIndicesWithMap(byteToRune, byteIndices[0], byteIndices[1])
				expectedRuneIndices[i] = []int{runeStart, runeEnd}
			}
			ass.ElementsMatch(expectedRuneIndices, findRegex2AllStringMatchIndex(word, re2))
		} else {
			ass.Nil(findRegex2AllStringMatchIndex(word, re2))
		}

		ass.Equal(re.FindString(word), findRegex2StringMatch(word, re2))
	}
}

func TestRegexWithUnicode(t *testing.T) {
	ass := assert.New(t)
	// Test with Unicode characters to verify byte-to-rune conversion works correctly
	pattern := `\p{L}+`
	re := regexp.MustCompile(pattern)
	re2 := regexp.MustCompile(pattern)

	testCases := []string{
		"hello世界",
		"Приветмир",
		"こんにちは世界",
		"hello 世界 world",
	}

	for _, text := range testCases {
		// Test FindStringIndex
		byteIndices := re.FindStringIndex(text)
		if byteIndices != nil {
			byteToRune := byteToRuneMap(text)
			runeStart, runeEnd := byteIndicesToRuneIndicesWithMap(byteToRune, byteIndices[0], byteIndices[1])
			expectedRuneIndices := []int{runeStart, runeEnd}
			actualRuneIndices := findRegex2StringIndex(text, re2)
			ass.ElementsMatch(expectedRuneIndices, actualRuneIndices, "Failed for text: %s", text)

			// Verify the matched substring is the same
			expectedMatch := text[byteIndices[0]:byteIndices[1]]
			actualMatch := string([]rune(text)[runeStart:runeEnd])
			ass.Equal(expectedMatch, actualMatch, "Matched substring should match for text: %s", text)
		}

		// Test FindAllStringIndex
		allByteIndices := re.FindAllStringIndex(text, -1)
		if allByteIndices != nil {
			byteToRune := byteToRuneMap(text)
			expectedRuneIndices := make([][]int, len(allByteIndices))
			for i, byteIndices := range allByteIndices {
				runeStart, runeEnd := byteIndicesToRuneIndicesWithMap(byteToRune, byteIndices[0], byteIndices[1])
				expectedRuneIndices[i] = []int{runeStart, runeEnd}
			}
			actualRuneIndices := findRegex2AllStringMatchIndex(text, re2)
			ass.ElementsMatch(expectedRuneIndices, actualRuneIndices, "Failed for text: %s", text)
		}
	}
}
