package simhash

import (
	"hash/fnv"
)

// SimHash:
// https://en.wikipedia.org/wiki/SimHash
// http://matpalm.com/resemblance/simhash/

const VectorSize = 64

type Vector [VectorSize]int

// type Feature interface {
// 	Sum64() uint64
// 	Weight() int
// }

type Interface interface {
	Len() int           // number of features
	Bytes(i int) []byte // bytes of i's feature
	Weight(i int) int
}

func Simhash(x Interface) uint64 {
	features := CalcFeatures(x)
	v := Vectorize(features)
	return Fingerprint(v)
}

type Feature struct {
	Sum64  uint64
	Weight int
}

func CalcFeatures(x Interface) []Feature {
	features := make([]Feature, x.Len())
	h := fnv.New64()
	for i := range features {
		h.Reset()
		h.Write(x.Bytes(i))
		sum := h.Sum64()
		features[i] = Feature{
			Sum64:  sum,
			Weight: x.Weight(i),
		}
	}
	return features
}

func Vectorize(features []Feature) Vector {
	var v Vector
	for _, feature := range features {
		var (
			sum    = feature.Sum64
			weight = feature.Weight
		)
		for i := range v {
			bit := ((sum >> i) & 1)
			if bit == 1 {
				v[i] += weight
			} else {
				v[i] -= weight
			}
		}
	}
	return v
}

func Fingerprint(v Vector) uint64 {
	var f uint64
	for i := range v {
		if v[i] >= 0 {
			f |= (1 << i)
		}
	}
	return f
}

// Compare calculates the Hamming distance between two 64-bit integers
func Compare(a, b uint64) int {
	var c int
	v := a ^ b
	for c = 0; v != 0; c++ {
		v &= v - 1
	}
	return c
}
