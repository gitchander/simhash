package simhash

import (
	"bytes"
	"strings"
)

//------------------------------------------------------------------------------
type ShinglingRunes struct {
	Runes []rune
	K     int
}

var _ Features = ShinglingRunes{}

func (p ShinglingRunes) Len() int {
	if (p.K <= 0) || (len(p.Runes) < p.K) {
		return 0
	}
	return (len(p.Runes) - p.K + 1)
}

func (p ShinglingRunes) Bytes(i int) []byte {
	return encodeRunes(p.Runes[i : i+p.K])
}

func (p ShinglingRunes) Weight(i int) int { return 1 }

//------------------------------------------------------------------------------
type ShinglingStrings struct {
	Strings []string
	K       int
}

var _ Features = ShinglingStrings{}

func (p ShinglingStrings) Len() int {
	if (p.K <= 0) || (len(p.Strings) < p.K) {
		return 0
	}
	return (len(p.Strings) - p.K + 1)
}

func (p ShinglingStrings) Bytes(i int) []byte {
	return encodeStringsV2(p.Strings[i : i+p.K])
}

func (p ShinglingStrings) Weight(i int) int { return 1 }

//------------------------------------------------------------------------------
func ShingleBytes(a [][]byte, k int) [][]byte {
	if k < 1 {
		panic("shingle has't elements")
	}
	if k == 1 {
		return a
	}
	n := len(a) - k + 1
	shingles := make([][]byte, n)
	for i := range shingles {
		shingles[i] = bytes.Join(a[i:i+k], []byte(" "))
	}
	return shingles
}

func ShingleStrings(a []string, k int) []string {
	if k < 1 {
		panic("shingle has't elements")
	}
	if k == 1 {
		return a
	}
	n := len(a) - k + 1
	shingles := make([]string, n)
	for i := range shingles {
		shingles[i] = strings.Join(a[i:i+k], " ")
	}
	return shingles
}
