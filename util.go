package simhash

import (
	// "fmt"
	"hash/fnv"
	"strings"
	"unicode/utf8"
)

type ByteSlice [][]byte

var _ Interface = ByteSlice{}

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Bytes(i int) []byte { return p[i] }
func (p ByteSlice) Weight(i int) int   { return 1 }

//-----------------------------------------------------------------------------
type Strings []string

var _ Interface = Strings{}

func (p Strings) Len() int           { return len(p) }
func (p Strings) Bytes(i int) []byte { return []byte(p[i]) }
func (p Strings) Weight(i int) int   { return 1 }

//-----------------------------------------------------------------------------
type ByFrase struct {
	Runes []rune
	N     int
}

var _ Interface = ByFrase{}

func (p ByFrase) Len() int {
	n := (len(p.Runes) - p.N + 1)
	if n < 0 {
		n = 0
	}
	return n
}

func (p ByFrase) Bytes(i int) []byte {
	data := make([]byte, (p.N * utf8.UTFMax))
	var k int
	for j := 0; j < p.N; j++ {
		n := utf8.EncodeRune(data[k:], p.Runes[i+j])
		k += n
	}
	//fmt.Printf("> %q\n", string(data[:k]))
	return data[:k]
}

func (p ByFrase) Weight(i int) int { return 1 }

//-----------------------------------------------------------------------------
// func getBit_v1(x uint64, i int) bool {
// 	return ((x & (1 << i)) != 0)
// }

// func getBit_v2(x uint64, i int) bool {
// 	return (((x >> i) & 1) == 1)
// }

func wordsFeatures(s string) []Feature {
	gs := strings.Split(s, " ")
	features := make([]Feature, len(gs))
	h := fnv.New64()
	for i := range features {
		h.Reset()
		h.Write([]byte(gs[i]))
		sum := h.Sum64()
		features[i] = Feature{
			Sum64:  sum,
			Weight: 1,
		}
	}
	return features
}

func stringFeatures(s string) []Feature {
	return runesFeatures([]rune(s))
}

func runesFeatures(rs []rune) []Feature {
	if len(rs) == 0 {
		return nil
	}
	h := fnv.New64()
	data := make([]byte, utf8.UTFMax)
	features := make([]Feature, len(rs)-1)
	for i := range features {

		h.Reset()

		n := utf8.EncodeRune(data, rs[i])
		h.Write(data[:n])

		n = utf8.EncodeRune(data, rs[i+1])
		h.Write(data[:n])

		features[i] = Feature{
			Sum64:  h.Sum64(),
			Weight: 1,
		}
	}
	return features
}

// func bytesFeatures(bs []byte) []Feature {
// 	if len(bs) == 0 {
// 		return nil
// 	}
// 	h := fnv.New64()
// 	features := make([]feature, len(bs)-1)
// 	for i := range features {
// 		h.Reset()
// 		h.Write(bs[i : i+2])
// 		features[i] = feature{
// 			sum:    h.Sum64(),
// 			weight: 1,
// 		}
// 	}
// 	return features
// }

// func bytesFeatures(bs [][]byte) []Feature {
// 	features := make([]Feature, len(bs))
// 	h := fnv.New64()
// 	for i := range features {
// 		h.Reset()
// 		h.Write(bs[i])
// 		sum := h.Sum64()
// 		features[i] = Feature{
// 			Sum64:  sum,
// 			Weight: 1,
// 		}
// 	}
// 	return features
// }
