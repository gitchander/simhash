package simhash

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func encodeRune(r rune) []byte {
	data := make([]byte, utf8.UTFMax)
	n := utf8.EncodeRune(data, r)
	return data[:n]
}

func encodeRunes(rs []rune) []byte {
	data := make([]byte, (len(rs) * utf8.UTFMax))
	var n int
	for _, r := range rs {
		n += utf8.EncodeRune(data[n:], r)
	}
	return data[:n]
}

func encodeStringsV1(vs []string) []byte {
	var n int
	for _, v := range vs {
		n += len(v)
	}
	data := make([]byte, n)
	var k int
	for _, v := range vs {
		copy(data[k:], v)
		k += len(v)
	}
	return data
}

func encodeStringsV2(as []string) []byte {
	var b bytes.Buffer
	for i, a := range as {
		if i > 0 {
			b.WriteByte('-') // string separator
		}
		b.WriteString(a)
	}
	return b.Bytes()
}

var punctuation = mapByRunes([]rune(".,;:-!?"))

func mapByRunes(rs []rune) map[rune]struct{} {
	m := make(map[rune]struct{})
	for _, r := range rs {
		m[r] = struct{}{}
	}
	return m
}

func runeIsPunctuation(r rune) bool {
	_, ok := punctuation[r]
	return ok
}

func removePunctuation(r rune) rune {
	if _, ok := punctuation[r]; ok {
		return -1
	}
	return r
}

// Words splits the string words
func Words(text string) []string {
	// remove punctuation
	s := strings.Map(removePunctuation, text)
	s = strings.ToLower(s)
	return strings.Fields(s)
}
