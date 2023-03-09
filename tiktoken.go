package tiktoken

import (
	"fmt"
	"regexp"
	"strings"
)

func GetEncoding(encodingName string) (*Tiktoken, error) {
	enc, err := getEncoding(encodingName)
	if err != nil {
		return nil, err
	}
	pbe, err := NewCoreBPE(enc.MergeableRanks, enc.SpecialTokens, enc.PatStr)
	if err != nil {
		return nil, err
	}
	specialTokensSet := map[string]bool{}
	for k := range enc.SpecialTokens {
		specialTokensSet[k] = true
	}
	return &Tiktoken{
		bpe:         pbe,
		pbeEncoding: enc,
	}, nil
}

func EncodingForModel(modelName string) (*Tiktoken, error) {
	if encodingName, ok := MODEL_TO_ENCODING[modelName]; !ok {
		return nil, fmt.Errorf("no encoding for model %s", modelName)
	} else {
		return GetEncoding(encodingName)
	}
}

type Tiktoken struct {
	bpe              *CoreBPE
	pbeEncoding      *Encoding
	specialTokensSet map[string]bool
}

func (t *Tiktoken) Encode(text string, allowedSpecial interface{}, disallowedSpecial interface{}) []int {
	var allowedSpecialSet map[string]any
	switch v := allowedSpecial.(type) {
	case string:
		if v == "all" {
			allowedSpecialSet = make(map[string]any, len(t.specialTokensSet))
			for token := range t.specialTokensSet {
				allowedSpecialSet[token] = true
			}
		}
	case map[string]any:
		allowedSpecialSet = v
	}

	var disallowedSpecialSet map[string]bool
	switch v := disallowedSpecial.(type) {
	case string:
		if v == "all" {
			disallowedSpecialSet = make(map[string]bool, len(t.specialTokensSet))
			for token := range t.specialTokensSet {
				disallowedSpecialSet[token] = true
			}
		}
	case map[string]bool:
		disallowedSpecialSet = v
	}

	if disallowedSpecialSet != nil {
		specialTokenRegex := t.SpecialTokenRegex(disallowedSpecialSet)
		match := specialTokenRegex.FindString(text)
		if match != "" {
			panic(fmt.Sprintf("Disallowed special token found: %v", match))
		}
	}

	tokens, _ := t.bpe.encodeNative(text, allowedSpecialSet)
	return tokens
}

func (t *Tiktoken) SpecialTokenRegex(disallowedSpecialSet map[string]bool) *regexp.Regexp {
	specialRegexStrs := make([]string, 0, len(disallowedSpecialSet))
	for k := range disallowedSpecialSet {
		specialRegexStrs = append(specialRegexStrs, regexp.QuoteMeta(k))
	}
	specialRegex, _ := regexp.Compile(strings.Join(specialRegexStrs, "|"))
	return specialRegex
}
