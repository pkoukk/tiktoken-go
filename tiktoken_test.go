package tiktoken

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := EncodingForModel("gpt-3.5-turbo-16k")
	ass.Nil(err, "Encoding  init should not be nil")
	tokens := enc.Encode("hello world!你好，世界！", []string{"all"}, []string{"all"})
	// these tokens are converted from the original python code
	sourceTokens := []int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, nil)
	sourceTokens = []int{15339, 220, 100257}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, []string{"all"})
	sourceTokens = []int{15339, 220, 100257}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	ass.Panics(func() {
		tokens = enc.Encode("hello <|endoftext|><|endofprompt|>", []string{"<|endoftext|>"}, []string{"all"})
	})
	ass.Panics(func() {
		tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, []string{"<|endoftext|>"})
	})
}

func TestDecoding(t *testing.T) {
	ass := assert.New(t)
	// enc, err := GetEncoding("cl100k_base")
	enc, err := GetEncoding(MODEL_CL100K_BASE)
	ass.Nil(err, "Encoding  init should not be nil")
	sourceTokens := []int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
	txt := enc.Decode(sourceTokens)
	ass.Equal("hello world!你好，世界！", txt, "Decoding should be equal")
}

func TestO200Decoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := GetEncoding(MODEL_O200K_BASE)
	ass.Nil(err, "Encoding init should not be nil")
	sourceTokens := []int{83, 8251, 2488, 382, 2212, 30}
	txt := enc.Decode(sourceTokens)
	ass.Equal("tiktoken is great?", txt, "Decoding should be equal")
}
