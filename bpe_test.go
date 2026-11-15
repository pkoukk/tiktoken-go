package tiktoken

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// uncachedBPE builds an encoding without adding it to the package cache, which
// TestGetEncoding_ErrorResponseNotCached expects not to hold o200k_base.
func uncachedBPE(t *testing.T, name string) *CoreBPE {
	t.Helper()
	enc, err := initEncoding(name)
	require.NoError(t, err)
	bpe, err := NewCoreBPE(enc.MergeableRanks, enc.SpecialTokens, enc.PatStr)
	require.NoError(t, err)
	return bpe
}

func TestBytePairMergeLargeMatchesScan(t *testing.T) {
	units := []string{"a", "ab", " ", "-", "1", "中", "😀", "é", "\n", "the ", "ing"}
	span := func(start, end int) [2]int { return [2]int{start, end} }
	for _, name := range []string{MODEL_CL100K_BASE, MODEL_O200K_BASE} {
		ranks := uncachedBPE(t, name).encoder
		random := rand.New(rand.NewSource(1))
		for i := 0; i < 300; i++ {
			length := 2 + random.Intn(3000)
			var piece strings.Builder
			for piece.Len() < length {
				piece.WriteString(strings.Repeat(units[random.Intn(len(units))], 1+random.Intn(8)))
			}
			raw := []byte(piece.String())
			assert.Equal(t, bytePairMergeScan(raw, ranks, span), bytePairMergeLarge(raw, ranks, span), "%s piece %q", name, piece.String())
		}
	}
}

func TestEncodeLongPiece(t *testing.T) {
	// these counts are from the original python code
	tests := []struct {
		unit   string
		repeat int
		cl100k int
		o200k  int
	}{
		{"a", 1000, 125, 125},
		{" ", 1000, 9, 9},
		{"中", 500, 500, 500},
		{"-", 1000, 16, 16},
		{"ab", 600, 600, 300},
		{"😀", 300, 600, 300},
	}
	cl100k, o200k := uncachedBPE(t, MODEL_CL100K_BASE), uncachedBPE(t, MODEL_O200K_BASE)
	for _, test := range tests {
		text := strings.Repeat(test.unit, test.repeat)
		assert.Len(t, cl100k.encodeOrdinaryNative(text), test.cl100k, "cl100k_base %d × %q", test.repeat, test.unit)
		assert.Len(t, o200k.encodeOrdinaryNative(text), test.o200k, "o200k_base %d × %q", test.repeat, test.unit)
	}
}

func BenchmarkEncodeLongPiece(b *testing.B) {
	enc, err := GetEncoding(MODEL_CL100K_BASE)
	require.NoError(b, err)
	text := strings.Repeat("中", 20000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.EncodeOrdinary(text)
	}
}
