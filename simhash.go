package simhash

import (
	"hash/fnv"
)

// SimHash:

type Vector [64]int

type Features interface {
	Len() int           // number of features
	Bytes(i int) []byte // bytes of i-th feature
	Weight(i int) int   // weight of i-th feature
}

func Simhash(fs Features) uint64 {
	v := Vectorize(fs)
	return Fingerprint(v)
}

func Vectorize(fs Features) Vector {
	var v Vector
	h := fnv.New64()
	n := fs.Len()
	for i := 0; i < n; i++ {
		h.Reset()
		h.Write(fs.Bytes(i))
		var (
			sum    = h.Sum64()
			weight = fs.Weight(i)
		)
		for j := range v {
			bit := ((sum >> j) & 1)
			if bit == 1 {
				v[j] += weight
			} else {
				v[j] -= weight
			}
		}
	}
	return v
}

func VectorizeBytes(bs [][]byte) Vector {
	var v Vector
	h := fnv.New64()
	for _, b := range bs {
		h.Reset()
		h.Write(b)
		var (
			sum    = h.Sum64()
			weight = 1
		)
		for j := range v {
			bit := ((sum >> j) & 1)
			if bit == 1 {
				v[j] += weight
			} else {
				v[j] -= weight
			}
		}
	}
	return v
}

func Fingerprint(v Vector) uint64 {
	var fp uint64
	for i := range v {
		if v[i] >= 0 {
			fp |= (1 << i)
		}
	}
	return fp
}

// Compare calculates the Hamming distance between two 64-bit integers
func Compare(a, b uint64) int {
	var n int
	v := a ^ b
	for n = 0; v != 0; n++ {
		v &= v - 1
	}
	return n
}
