package tiktoken

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := GetEncoding("cl100k_base")
	ass.Nil(err, "Encoding  init should not be nil")
	tokens := enc.Encode("hello world!你好，世界！", nil, nil)
	// these tokens are converted from the original python code
	ass.Equal([]int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}, tokens, "Encoding should be equal")
}

func TestDecoding(t *testing.T) {

}
