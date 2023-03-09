package tiktoken

import "math"

func bytePairMerge[T any](piece []byte, ranks map[string]int, f func(start, end int) T) []T {
	parts := make([][2]int, len(piece)+1)
	for i := 0; i < len(parts); i++ {
		parts[i][0], parts[i][1] = i, math.MaxInt // use max int as sentinel
	}

	getRank := func(startIdx, skip int) int {
		if startIdx+skip+2 < len(parts) {
			rank, ok := ranks[string(piece[parts[startIdx][0]:parts[startIdx+skip+2][0]])]
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
			parts[i][1] = getRank(i, 1)
			if i > 0 {
				parts[i-1][1] = getRank(i-1, 1)
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

func bytePairDecode(piece []byte, ranks map[string]int) [][]byte {
	if len(piece) == 1 {
		return [][]byte{piece}
	}
	return bytePairMerge(piece, ranks, func(start, end int) []byte {
		return piece[start:end]
	})
}
