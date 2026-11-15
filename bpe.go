package tiktoken

import (
	"math"
)

// largePiece is the piece length from which bytePairMerge uses a heap, as in
// openai/tiktoken.
const largePiece = 100

func bytePairMerge[T any](piece []byte, ranks map[string]int, f func(start, end int) T) []T {
	if len(piece) >= largePiece {
		return bytePairMergeLarge(piece, ranks, f)
	}
	return bytePairMergeScan(piece, ranks, f)
}

// bytePairMergeScan repeatedly merges the leftmost pair with the lowest rank,
// finding it with a linear scan, so a piece of n bytes costs O(n^2).
func bytePairMergeScan[T any](piece []byte, ranks map[string]int, f func(start, end int) T) []T {
	parts := make([][2]int, len(piece)+1)
	for i := 0; i < len(parts); i++ {
		parts[i][0], parts[i][1] = i, math.MaxInt // use max int as sentinel
	}

	getRank := func(startIdx, skip int) int {
		if startIdx+skip+2 < len(parts) {
			b := piece[parts[startIdx][0]:parts[startIdx+skip+2][0]]
			rank, ok := ranks[string(b)]
			if ok {
				return rank
			}
		}
		return -1 // use -1 to represent None
	}

	for i := 0; i < len(parts)-2; i++ {
		if rank := getRank(i, 0); rank >= 0 {
			parts[i][1] = rank
		}
	}

	for len(parts) > 1 {
		minRank, minIdx := math.MaxInt, -1
		for i := 0; i < len(parts)-1; i++ {
			if parts[i][1] < minRank {
				minRank, minIdx = parts[i][1], i
			}
		}

		if minRank < math.MaxInt {
			i := minIdx
			rank := getRank(i, 1)
			if rank >= 0 {
				parts[i][1] = rank
			} else {
				parts[i][1] = math.MaxInt
			}
			if i > 0 {
				rk := getRank(i-1, 1)
				if rk >= 0 {
					parts[i-1][1] = rk
				} else {
					parts[i-1][1] = math.MaxInt
				}
			}
			parts = append(parts[:i+1], parts[i+2:]...)
		} else {
			break
		}
	}

	out := make([]T, len(parts)-1)
	for i := 0; i < len(out); i++ {
		out[i] = f(parts[i][0], parts[i+1][0])
	}
	return out
}

func bytePairEncode(piece []byte, ranks map[string]int) []int {
	if len(piece) == 1 {
		v := ranks[string(piece)]
		return []int{v}
	}
	return bytePairMerge(piece, ranks, func(start, end int) int {
		return ranks[string(piece[start:end])]
	})
}

// mergeState is the token that starts at a byte offset of the piece being
// merged, together with the merge of that token and the next one.
type mergeState struct {
	prev     int // start of the previous token
	end      int // end of this token
	nextEnd  int // end of the next token
	nextRank int // rank of this token merged with the next one, or math.MaxInt
}

// merge is a candidate merge of the token starting at start with the next one.
type merge struct {
	start int
	rank  int
}

// mergeHeap is a min-heap of candidate merges ordered by rank, then by start.
type mergeHeap []merge

func (h mergeHeap) less(i, j int) bool {
	return h[i].rank < h[j].rank || h[i].rank == h[j].rank && h[i].start < h[j].start
}

func (h *mergeHeap) push(m merge) {
	*h = append(*h, m)
	q := *h
	for i := len(q) - 1; i > 0; {
		parent := (i - 1) / 2
		if !q.less(i, parent) {
			break
		}
		q[i], q[parent] = q[parent], q[i]
		i = parent
	}
}

func (h *mergeHeap) pop() merge {
	q := *h
	top := q[0]
	last := len(q) - 1
	q[0] = q[last]
	q = q[:last]
	for i := 0; ; {
		smallest, left, right := i, 2*i+1, 2*i+2
		if left < len(q) && q.less(left, smallest) {
			smallest = left
		}
		if right < len(q) && q.less(right, smallest) {
			smallest = right
		}
		if smallest == i {
			break
		}
		q[i], q[smallest] = q[smallest], q[i]
		i = smallest
	}
	*h = q
	return top
}

// bytePairMergeLarge returns the same tokens as bytePairMergeScan. It keeps
// the candidate merges in a heap, so a piece of n bytes costs O(n log n)
// instead of O(n^2). It is a port of _byte_pair_merge_large in openai/tiktoken.
func bytePairMergeLarge[T any](piece []byte, ranks map[string]int, f func(start, end int) T) []T {
	states := make([]mergeState, len(piece))
	candidates := make(mergeHeap, 0, len(piece))

	// propose records the merge of the token at start with the token that ends
	// at nextEnd, replacing the previous candidate of that token.
	propose := func(start, nextEnd int) {
		states[start].nextEnd = nextEnd
		states[start].nextRank = math.MaxInt
		if nextEnd <= len(piece) {
			if rank, ok := ranks[string(piece[start:nextEnd])]; ok {
				states[start].nextRank = rank
				candidates.push(merge{start: start, rank: rank})
			}
		}
	}
	for i := range states {
		states[i] = mergeState{prev: i - 1, end: i + 1}
		propose(i, i+2)
	}

	tokens := len(piece)
	for len(candidates) > 0 {
		left := candidates.pop()
		if left.rank != states[left.start].nextRank {
			continue // replaced by a later candidate
		}
		rightStart := states[left.start].end
		rightEnd := states[left.start].nextEnd
		rightNextEnd := states[rightStart].nextEnd

		states[left.start].end = rightEnd
		propose(left.start, rightNextEnd)
		if rightEnd < len(piece) {
			states[rightEnd].prev = left.start
		}
		if left.start > 0 {
			propose(states[left.start].prev, rightEnd)
		}
		states[rightStart].nextRank = math.MaxInt
		tokens--
	}

	out := make([]T, 0, tokens)
	for start := 0; start < len(piece); start = states[start].end {
		out = append(out, f(start, states[start].end))
	}
	return out
}
