package simhash

// RunesGroup
type RunesGroup struct {
	Runes []rune
	N     int
}

var _ Features = RunesGroup{}

func (p RunesGroup) Len() int {
	if (p.N <= 0) || (len(p.Runes) < p.N) {
		return 0
	}
	return (len(p.Runes) - p.N + 1)
}

func (p RunesGroup) Bytes(i int) []byte {
	return encodeRunes(p.Runes[i : i+p.N])
}

func (p RunesGroup) Weight(i int) int { return 1 }

// StringsGroup
type StringsGroup struct {
	Strings []string
	N       int
}

var _ Features = StringsGroup{}

func (p StringsGroup) Len() int {
	if (p.N <= 0) || (len(p.Strings) < p.N) {
		return 0
	}
	return (len(p.Strings) - p.N + 1)
}

func (p StringsGroup) Bytes(i int) []byte {
	return encodeStringsV2(p.Strings[i : i+p.N])
}

func (p StringsGroup) Weight(i int) int { return 1 }
