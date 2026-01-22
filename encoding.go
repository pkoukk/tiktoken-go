package tiktoken

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

const ENDOFTEXT string = "<|endoftext|>"
const FIM_PREFIX string = "<|fim_prefix|>"
const FIM_MIDDLE string = "<|fim_middle|>"
const FIM_SUFFIX string = "<|fim_suffix|>"
const ENDOFPROMPT string = "<|endofprompt|>"
const STARTOFTEXT string = "<|startoftext|>"
const RETURN string = "<|return|>"
const CONSTRAIN string = "<|constrain|>"
const CHANNEL string = "<|channel|>"
const START string = "<|start|>"
const END string = "<|end|>"
const MESSAGE string = "<|message|>"
const CALL string = "<|call|>"

const (
	MODEL_O200K_BASE    string = "o200k_base"
	MODEL_O200K_HARMONY string = "o200k_harmony"
	MODEL_CL100K_BASE   string = "cl100k_base"
	MODEL_P50K_BASE     string = "p50k_base"
	MODEL_P50K_EDIT     string = "p50k_edit"
	MODEL_R50K_BASE     string = "r50k_base"
)

var MODEL_TO_ENCODING = map[string]string{
	// chat
	"gpt-5.2":       MODEL_O200K_BASE,
	"gpt-5.1":       MODEL_O200K_BASE,
	"gpt-5":         MODEL_O200K_BASE,
	"gpt-4.5":       MODEL_O200K_BASE,
	"gpt-4.1":       MODEL_O200K_BASE,
	"gpt-4o":        MODEL_O200K_BASE,
	"gpt-4":         MODEL_CL100K_BASE,
	"gpt-3.5-turbo": MODEL_CL100K_BASE,
	// o-series
	"o4-mini": MODEL_O200K_BASE,
	"o3-pro":  MODEL_O200K_BASE,
	"o3-mini": MODEL_O200K_BASE,
	"o3":      MODEL_O200K_BASE,
	"o1-pro":  MODEL_O200K_BASE,
	"o1":      MODEL_O200K_BASE,
	// text
	"text-davinci-003": MODEL_P50K_BASE,
	"text-davinci-002": MODEL_P50K_BASE,
	"text-davinci-001": MODEL_R50K_BASE,
	"davinci-002":      MODEL_CL100K_BASE,
	"babbage-002":      MODEL_CL100K_BASE,
	"text-curie-001":   MODEL_R50K_BASE,
	"text-babbage-001": MODEL_R50K_BASE,
	"text-ada-001":     MODEL_R50K_BASE,
	"davinci":          MODEL_R50K_BASE,
	"curie":            MODEL_R50K_BASE,
	"babbage":          MODEL_R50K_BASE,
	"ada":              MODEL_R50K_BASE,
	// code
	"code-davinci-002": MODEL_P50K_BASE,
	"code-davinci-001": MODEL_P50K_BASE,
	"code-cushman-002": MODEL_P50K_BASE,
	"code-cushman-001": MODEL_P50K_BASE,
	"davinci-codex":    MODEL_P50K_BASE,
	"cushman-codex":    MODEL_P50K_BASE,
	// edit
	"text-davinci-edit-001": MODEL_P50K_EDIT,
	"code-davinci-edit-001": MODEL_P50K_EDIT,
	// embeddings
	"text-embedding-ada-002": MODEL_CL100K_BASE,
	"text-embedding-3-large": MODEL_CL100K_BASE,
	"text-embedding-3-small": MODEL_CL100K_BASE,
	// old embeddings
	"text-similarity-davinci-001":  MODEL_R50K_BASE,
	"text-similarity-curie-001":    MODEL_R50K_BASE,
	"text-similarity-babbage-001":  MODEL_R50K_BASE,
	"text-similarity-ada-001":      MODEL_R50K_BASE,
	"text-search-davinci-doc-001":  MODEL_R50K_BASE,
	"text-search-curie-doc-001":    MODEL_R50K_BASE,
	"text-search-babbage-doc-001":  MODEL_R50K_BASE,
	"text-search-ada-doc-001":      MODEL_R50K_BASE,
	"code-search-babbage-code-001": MODEL_R50K_BASE,
	"code-search-ada-code-001":     MODEL_R50K_BASE,
	// open source
	"gpt2":         "gpt2",
	"gpt-oss-20b":  MODEL_O200K_HARMONY,
	"gpt-oss-120b": MODEL_O200K_HARMONY,
}

var MODEL_PREFIX_TO_ENCODING = map[string]string{
	// chat
	"gpt-5.2-":       MODEL_O200K_BASE,  // e.g., gpt-5.2-2025-12-11, etc.
	"gpt-5.1-":       MODEL_O200K_BASE,  // e.g., gpt-5.1-2025-11-13, etc.
	"gpt-5-":         MODEL_O200K_BASE,  // e.g., gpt-5-2025-08-07, etc.
	"gpt-4.5-":       MODEL_O200K_BASE,  // e.g., gpt-4.5-preview, etc.
	"gpt-4.1-":       MODEL_O200K_BASE,  // e.g., gpt-4.1-2025-04-14, etc.
	"gpt-4o-":        MODEL_O200K_BASE,  // e.g., gpt-4o-2024-05-13, etc.
	"gpt-4-":         MODEL_CL100K_BASE, // e.g., gpt-4-0314, etc., plus gpt-4-32k
	"gpt-3.5-turbo-": MODEL_CL100K_BASE, // e.g, gpt-3.5-turbo-0301, -0401, etc.
	// o-series
	"o4-mini-": MODEL_O200K_BASE, // e.g., o4-mini-2025-04-16, etc.
	"o3-":      MODEL_O200K_BASE, // e.g., o3-2025-04-16, etc.
	"o1-":      MODEL_O200K_BASE, // e.g., o1-2024-12-17, etc.
}

var encodingMap map[string]*Encoding
var l *sync.RWMutex

func init() {
	encodingMap = make(map[string]*Encoding)
	l = &sync.RWMutex{}
}

type Encoding struct {
	Name           string
	PatStr         string
	MergeableRanks map[string]int
	SpecialTokens  map[string]int
	ExplicitNVocab int
}

func getEncoding(encodingName string) (*Encoding, error) {
	l.RLock()
	encoding, ok := encodingMap[encodingName]
	l.RUnlock()
	if ok {
		return encoding, nil
	}

	l.Lock()
	defer l.Unlock()
	initEncoding, err := initEncoding(encodingName)
	if err != nil {
		return nil, err
	}
	encodingMap[encodingName] = initEncoding
	return encodingMap[encodingName], nil
}

func initEncoding(encodingName string) (*Encoding, error) {
	switch encodingName {
	case MODEL_O200K_BASE:
		return o200k_base()
	case MODEL_O200K_HARMONY:
		return o200k_harmony()
	case MODEL_CL100K_BASE:
		return cl100k_base()
	case MODEL_P50K_BASE:
		return p50k_base()
	case MODEL_R50K_BASE:
		return r50k_base()
	case MODEL_P50K_EDIT:
		return p50k_edit()
	default:
		return nil, errors.New("Unknown encoding: " + encodingName)
	}
}

func o200k_base() (*Encoding, error) {
	ranks, err := bpeLoader.LoadTiktokenBpe(
		"https://openaipublic.blob.core.windows.net/encodings/o200k_base.tiktoken",
	)
	if err != nil {
		return nil, err
	}
	special_tokens := map[string]int{
		ENDOFTEXT:   199999,
		ENDOFPROMPT: 200018,
	}
	pats := []string{
		`[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?`,
		`[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?`,
		`\p{N}{1,3}`,
		` ?[^\s\p{L}\p{N}]+[\r\n/]*`,
		`\s*[\r\n]+`,
		`\s+(?!\S)`,
		`\s+`,
	}
	return &Encoding{
		Name:           MODEL_O200K_BASE,
		PatStr:         strings.Join(pats, "|"),
		MergeableRanks: ranks,
		SpecialTokens:  special_tokens,
	}, nil
}

func o200k_harmony() (*Encoding, error) {
	// Based on o200k_base
	base_enc, err := o200k_base()
	if err != nil {
		return nil, err
	}

	// Start with base special tokens and add harmony-specific ones
	special_tokens := make(map[string]int)
	for k, v := range base_enc.SpecialTokens {
		special_tokens[k] = v
	}

	// Add o200k_harmony specific tokens
	special_tokens[STARTOFTEXT] = 199998
	special_tokens[ENDOFTEXT] = 199999
	special_tokens["<|reserved_200000|>"] = 200000
	special_tokens["<|reserved_200001|>"] = 200001
	special_tokens[RETURN] = 200002
	special_tokens[CONSTRAIN] = 200003
	special_tokens["<|reserved_200004|>"] = 200004
	special_tokens[CHANNEL] = 200005
	special_tokens[START] = 200006
	special_tokens[END] = 200007
	special_tokens[MESSAGE] = 200008
	special_tokens["<|reserved_200009|>"] = 200009
	special_tokens["<|reserved_200010|>"] = 200010
	special_tokens["<|reserved_200011|>"] = 200011
	special_tokens[CALL] = 200012

	// Add reserved tokens from 200013 to 201087
	for i := 200013; i < 201088; i++ {
		special_tokens[fmt.Sprintf("<|reserved_%d|>", i)] = i
	}

	return &Encoding{
		Name:           MODEL_O200K_HARMONY,
		PatStr:         base_enc.PatStr,
		MergeableRanks: base_enc.MergeableRanks,
		SpecialTokens:  special_tokens,
	}, nil
}

func cl100k_base() (*Encoding, error) {
	ranks, err := bpeLoader.LoadTiktokenBpe(
		"https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken",
	)
	if err != nil {
		return nil, err
	}
	special_tokens := map[string]int{
		ENDOFTEXT:   100257,
		FIM_PREFIX:  100258,
		FIM_MIDDLE:  100259,
		FIM_SUFFIX:  100260,
		ENDOFPROMPT: 100276,
	}
	return &Encoding{
		Name:           MODEL_CL100K_BASE,
		PatStr:         `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`,
		MergeableRanks: ranks,
		SpecialTokens:  special_tokens,
	}, nil
}

func p50k_edit() (*Encoding, error) {
	ranks, err := bpeLoader.LoadTiktokenBpe(
		"https://openaipublic.blob.core.windows.net/encodings/p50k_base.tiktoken",
	)
	if err != nil {
		return nil, err
	}
	special_tokens := map[string]int{
		ENDOFTEXT:  50256,
		FIM_PREFIX: 50281,
		FIM_MIDDLE: 50282,
		FIM_SUFFIX: 50283,
	}
	return &Encoding{
		Name:           MODEL_P50K_EDIT,
		PatStr:         `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`,
		MergeableRanks: ranks,
		SpecialTokens:  special_tokens,
	}, nil
}

func p50k_base() (*Encoding, error) {
	ranks, err := bpeLoader.LoadTiktokenBpe(
		"https://openaipublic.blob.core.windows.net/encodings/p50k_base.tiktoken",
	)
	if err != nil {
		return nil, err
	}
	special_tokens := map[string]int{ENDOFTEXT: 50256}

	// ExplicitNVocab := 50281
	// max_tokens := int(math.Max(float64(len(special_tokens)), float64(len(ranks))))

	// if len(special_tokens)+len(ranks) != max_tokens {
	// 	return nil, errors.New("special_tokens and ranks must be disjoint")
	// }

	return &Encoding{
		Name:           MODEL_P50K_BASE,
		PatStr:         `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`,
		MergeableRanks: ranks,
		SpecialTokens:  special_tokens,
		ExplicitNVocab: 50281,
	}, nil
}

func r50k_base() (*Encoding, error) {
	ranks, err := bpeLoader.LoadTiktokenBpe(
		"https://openaipublic.blob.core.windows.net/encodings/r50k_base.tiktoken",
	)
	if err != nil {
		return nil, err
	}
	special_tokens := map[string]int{ENDOFTEXT: 50256}
	return &Encoding{
		Name:           MODEL_R50K_BASE,
		MergeableRanks: ranks,
		PatStr:         `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`,
		SpecialTokens:  special_tokens,
		ExplicitNVocab: 50257,
	}, nil
}

// var ENCODING_MAP = map[string]*Encoding{}
