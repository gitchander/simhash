package simhash

import (
	"strings"
)

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
	return strings.Fields(s)
}
