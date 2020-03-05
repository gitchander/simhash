package simhash

import (
	"unicode/utf8"
)

//-----------------------------------------------------------------------------
type BytesSlice [][]byte

var _ Features = BytesSlice{}

func (p BytesSlice) Len() int           { return len(p) }
func (p BytesSlice) Bytes(i int) []byte { return p[i] }
func (p BytesSlice) Weight(i int) int   { return 1 }

//-----------------------------------------------------------------------------
type StringSlice []string

var _ Features = StringSlice{}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Bytes(i int) []byte { return []byte(p[i]) }
func (p StringSlice) Weight(i int) int   { return 1 }

//-----------------------------------------------------------------------------
type RuneSlice []rune

var _ Features = RuneSlice{}

func (p RuneSlice) Len() int { return len(p) }

func (p RuneSlice) Bytes(i int) []byte {
	data := make([]byte, utf8.UTFMax)
	n := utf8.EncodeRune(data, p[i])
	return data[:n]
}

func (p RuneSlice) Weight(i int) int { return 1 }

//-----------------------------------------------------------------------------
type RunesGroup struct {
	Runes []rune
	N     int
}

var _ Features = RunesGroup{}

func (p RunesGroup) Len() int {
	n := (len(p.Runes) - p.N + 1)
	if n < 0 {
		n = 0
	}
	return n
}

func (p RunesGroup) Bytes(i int) []byte {
	data := make([]byte, (p.N * utf8.UTFMax))
	var k int
	for j := 0; j < p.N; j++ {
		n := utf8.EncodeRune(data[k:], p.Runes[i+j])
		k += n
	}
	//fmt.Printf("> %q\n", string(data[:k]))
	return data[:k]
}

func (p RunesGroup) Weight(i int) int { return 1 }

//-----------------------------------------------------------------------------
type StringsGroup struct {
	Strings []string
	N       int
}

var _ Features = StringsGroup{}

func (p StringsGroup) Len() int {
	n := (len(p.Strings) - p.N + 1)
	if n < 0 {
		n = 0
	}
	return n
}

func (p StringsGroup) Bytes(i int) []byte {
	var data []byte
	for j := 0; j < p.N; j++ {
		s := p.Strings[i+j]
		data = append(data, []byte(s)...)
	}
	//fmt.Printf("> %q\n", string(data))
	return data
}

func (p StringsGroup) Weight(i int) int { return 1 }

//-----------------------------------------------------------------------------
