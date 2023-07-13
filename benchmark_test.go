package tiktoken

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const TEST_FILE = "test/test.txt"

func ReadTestFile() ([]byte, error) {
	// open and read TEST_FILE
	return os.ReadFile(TEST_FILE)
}

func BenchmarkEncoding(b *testing.B) {
	fileContent, err := ReadTestFile()
	if err != nil {
		panic(err)
	}

	tkm, err := EncodingForModel("gpt-4")
	if err != nil {
		panic(err)
	}

	text := string(fileContent)

	for ordersOfMagnitude := 0; ordersOfMagnitude < 4; ordersOfMagnitude++ {
		// do actual encoding
		fmt.Printf("Encoding %d bytes\n", len(text))
		tkm.Encode(text, nil, nil)

		stringBuilder := strings.Builder{}
		for i := 0; i < 10; i++ {
			stringBuilder.WriteString(text)
		}

		text = stringBuilder.String()
	}

}
