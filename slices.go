package simhash

// BytesSlice
type BytesSlice [][]byte

var _ Features = BytesSlice{}

func (p BytesSlice) Len() int           { return len(p) }
func (p BytesSlice) Bytes(i int) []byte { return p[i] }
func (p BytesSlice) Weight(i int) int   { return 1 }

// StringSlice
type StringSlice []string

var _ Features = StringSlice{}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Bytes(i int) []byte { return []byte(p[i]) }
func (p StringSlice) Weight(i int) int   { return 1 }

// RuneSlice
type RuneSlice []rune

var _ Features = RuneSlice{}

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Bytes(i int) []byte { return encodeRune(p[i]) }
func (p RuneSlice) Weight(i int) int   { return 1 }
