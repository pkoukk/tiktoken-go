package tiktoken

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const TEST_FILE = "test/test.txt"

var (
	testTexts = map[string]string{
		"small":   "Hello, world!",
		"medium":  strings.Repeat("Hello, world! This is a medium length text. ", 100),
		"large":   strings.Repeat("Hello, world! This is a large text to test performance. ", 1000),
		"unicode": "你好世界！こんにちは世界！안녕하세요 세계! Привет мир! ¡Hola mundo!",
		"code":    strings.Repeat("func main() { fmt.Println(\"Hello, world!\") }\n", 50),
		"mixed":   "Hello 世界! func main() { fmt.Println(\"你好\") } こんにちは!",
		"special": "Hello <|endoftext|> world!",
	}
	testModels = []string{
		"gpt-4o",
		"gpt-4",
		"gpt-3.5-turbo",
		"text-davinci-003",
		"text-davinci-002",
		"text-embedding-ada-002",
	}
)

func ReadTestFile() ([]byte, error) {
	return os.ReadFile(TEST_FILE)
}

func getTestText(name string) (string, error) {
	if text, ok := testTexts[name]; ok {
		return text, nil
	}
	fileContent, err := ReadTestFile()
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

// Benchmark different models encoding
func BenchmarkEncoding_Models(b *testing.B) {
	for _, model := range testModels {
		b.Run(model, func(b *testing.B) {
			tkm, err := EncodingForModel(model)
			if err != nil {
				b.Fatalf("Failed to get encoding for %s: %v", model, err)
			}
			text, err := getTestText("medium")
			if err != nil {
				b.Fatalf("Failed to get test text: %v", err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tkm.Encode(text, nil, nil)
			}
		})
	}
}

// Benchmark different text sizes
func BenchmarkEncoding_TextSizes(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	sizes := []string{"small", "medium", "large"}
	for _, size := range sizes {
		b.Run(size, func(b *testing.B) {
			text, err := getTestText(size)
			if err != nil {
				b.Fatalf("Failed to get test text: %v", err)
			}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tkm.Encode(text, nil, nil)
			}
		})
	}
}

// Benchmark different text types
func BenchmarkEncoding_TextTypes(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	textTypes := []string{"small", "unicode", "code", "mixed"}
	for _, textType := range textTypes {
		b.Run(textType, func(b *testing.B) {
			text, err := getTestText(textType)
			if err != nil {
				b.Fatalf("Failed to get test text: %v", err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tkm.Encode(text, nil, nil)
			}
		})
	}
}

// Benchmark EncodeOrdinary vs Encode
func BenchmarkEncoding_EncodeOrdinary(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text, err := getTestText("medium")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	b.Run("Encode", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tkm.Encode(text, nil, nil)
		}
	})

	b.Run("EncodeOrdinary", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tkm.EncodeOrdinary(text)
		}
	})
}

// Benchmark decoding
func BenchmarkDecoding(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text, err := getTestText("medium")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	tokens := tkm.Encode(text, nil, nil)

	b.Run("Decode", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			tkm.Decode(tokens)
		}
	})

	// Test with different token list sizes
	tokenSizes := []int{10, 100, 1000}
	for _, size := range tokenSizes {
		if size > len(tokens) {
			continue
		}
		tokenSlice := tokens[:size]
		b.Run(fmt.Sprintf("Decode_%d_tokens", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tkm.Decode(tokenSlice)
			}
		})
	}
}

// Benchmark round-trip encoding/decoding
func BenchmarkRoundTrip(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	textTypes := []string{"small", "medium", "unicode", "mixed"}
	for _, textType := range textTypes {
		b.Run(textType, func(b *testing.B) {
			text, err := getTestText(textType)
			if err != nil {
				b.Fatalf("Failed to get test text: %v", err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tokens := tkm.Encode(text, nil, nil)
				decoded := tkm.Decode(tokens)
				_ = decoded // Prevent optimization
			}
		})
	}
}

// Benchmark special token handling
func BenchmarkSpecialTokens(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text, err := getTestText("special")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}

	b.Run("WithAllowedSpecial", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tkm.Encode(text, []string{"<|endoftext|>"}, nil)
		}
	})

	b.Run("WithAllAllowedSpecial", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tkm.Encode(text, []string{"all"}, nil)
		}
	})

	b.Run("EncodeOrdinary_IgnoresSpecial", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tkm.EncodeOrdinary(text)
		}
	})
}

// Benchmark concurrent encoding
func BenchmarkConcurrentEncoding(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text, err := getTestText("medium")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tkm.Encode(text, nil, nil)
		}
	})
}

// Benchmark encoding with different input sizes (scaling)
func BenchmarkEncoding_Scaling(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	baseText, err := getTestText("small")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	for ordersOfMagnitude := 0; ordersOfMagnitude < 5; ordersOfMagnitude++ {
		text := baseText
		for i := 0; i < ordersOfMagnitude; i++ {
			text = strings.Repeat(text, 10)
		}

		b.Run(fmt.Sprintf("Size_10^%d", ordersOfMagnitude), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tkm.Encode(text, nil, nil)
			}
		})
	}
}

// Benchmark memory allocation for encoding
func BenchmarkEncoding_Allocations(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	textTypes := []string{"small", "medium", "large"}
	for _, textType := range textTypes {
		b.Run(textType, func(b *testing.B) {
			text, err := getTestText(textType)
			if err != nil {
				b.Fatalf("Failed to get test text: %v", err)
			}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tokens := tkm.Encode(text, nil, nil)
				_ = tokens // Prevent optimization
			}
		})
	}
}

// Benchmark file-based encoding (original test)
func BenchmarkEncoding_File(b *testing.B) {
	fileContent, err := ReadTestFile()
	if err != nil {
		b.Fatalf("Failed to read test file: %v", err)
	}

	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text := string(fileContent)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tkm.Encode(text, nil, nil)
	}
}

// Benchmark multiple encodings initialization
func BenchmarkEncoding_Initialization(b *testing.B) {
	for _, model := range testModels {
		b.Run(model, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tkm, err := EncodingForModel(model)
				if err != nil {
					b.Fatalf("Failed to get encoding for %s: %v", model, err)
				}
				_ = tkm // Prevent optimization
			}
		})
	}
}

// Benchmark encoding with various special token configurations
func BenchmarkEncoding_SpecialTokenConfigs(b *testing.B) {
	tkm, err := EncodingForModel("gpt-3.5-turbo")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text := "Hello <|endoftext|> world <|fim_prefix|> test"

	configs := []struct {
		name              string
		allowedSpecial    []string
		disallowedSpecial []string
	}{
		{"NoSpecial", nil, nil},
		{"AllowEndOfText", []string{"<|endoftext|>"}, nil},
		{"AllowAll", []string{"all"}, nil},
		{"DisallowAll", nil, []string{"all"}},
		{"AllowOneDisallowOther", []string{"<|endoftext|>"}, []string{"all"}},
	}

	for _, config := range configs {
		b.Run(config.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Use recover to handle panics from disallowed special tokens
				func() {
					defer func() {
						if r := recover(); r != nil {
							// Expected panic for disallowed special tokens
						}
					}()
					tkm.Encode(text, config.allowedSpecial, config.disallowedSpecial)
				}()
			}
		})
	}
}

// Benchmark encoding across different encodings (not models)
func BenchmarkEncoding_Encodings(b *testing.B) {
	encodings := []string{
		MODEL_O200K_BASE,
		MODEL_CL100K_BASE,
		MODEL_P50K_BASE,
		MODEL_R50K_BASE,
	}

	text, err := getTestText("medium")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	for _, encoding := range encodings {
		b.Run(encoding, func(b *testing.B) {
			tkm, err := GetEncoding(encoding)
			if err != nil {
				b.Fatalf("Failed to get encoding %s: %v", encoding, err)
			}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tkm.Encode(text, nil, nil)
			}
		})
	}
}

// Benchmark large batch encoding
func BenchmarkEncoding_LargeBatch(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text, err := getTestText("medium")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	texts := make([]string, 100)
	for i := range texts {
		texts[i] = text
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, text := range texts {
			tokens := tkm.Encode(text, nil, nil)
			_ = tokens
		}
	}
}

// Benchmark concurrent encoding and decoding
func BenchmarkConcurrent_EncodeDecode(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	text, err := getTestText("medium")
	if err != nil {
		b.Fatalf("Failed to get test text: %v", err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tokens := tkm.Encode(text, nil, nil)
			decoded := tkm.Decode(tokens)
			_ = decoded
		}
	})
}

// Benchmark with realistic text patterns
func BenchmarkEncoding_RealisticPatterns(b *testing.B) {
	tkm, err := EncodingForModel("gpt-4o")
	if err != nil {
		b.Fatalf("Failed to get encoding: %v", err)
	}

	patterns := map[string]string{
		"chat":       "User: Hello! How are you?\nAssistant: I'm doing well, thank you!",
		"code":       "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, world!\")\n}",
		"json":       `{"name": "test", "value": 123, "nested": {"key": "value"}}`,
		"markdown":   "# Title\n\nThis is a **bold** text with *italic* and `code`.",
		"multiline":  strings.Repeat("Line with some text.\n", 50),
		"whitespace": strings.Repeat("   \t\n  ", 100),
	}

	for name, text := range patterns {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tkm.Encode(text, nil, nil)
			}
		})
	}
}
