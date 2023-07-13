package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/pkoukk/tiktoken-go"
)

func BenchmarkEncodingInFullLanguage(b *testing.B) {
	// Universal Declaration of Human Rights in all languages
	url := "https://unicode.org/udhr/assemblies/full_all.txt"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	lines := strings.Split(responseString, "\n")
	tkm, err := tiktoken.EncodingForModel("gpt-4")
	lineCount := len(lines)
	if err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tkm.EncodeOrdinary(lines[n%lineCount])
	}
}
