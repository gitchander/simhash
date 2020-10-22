package simhash

import (
	"hash"
	"hash/fnv"
)

type Feature struct {
	Hash   uint64
	Weight int
}

func VectorizeFeatures(fs []Feature) Vector {
	var v Vector
	for _, f := range fs {
		for i := range v {
			bit := getBit(f.Hash, i)
			switch bit {
			case 0:
				v[i] -= f.Weight
			case 1:
				v[i] += f.Weight
			}
		}
	}
	return v
}

func getBit(x uint64, i int) uint {
	return uint((x >> i) & 1)
}

func FeaturesByStrings(vs []string) []Feature {

	var h hash.Hash64
	h = fnv.New64()
	weight := 1

	fs := make([]Feature, len(vs))
	for i, v := range vs {
		h.Reset()
		h.Write([]byte(v))
		sum := h.Sum64()

		fs[i] = Feature{
			Hash:   sum,
			Weight: weight,
		}
	}
	return fs
}

func FeaturesByBytesSlices(vs [][]byte) []Feature {

	var h hash.Hash64
	h = fnv.New64()
	weight := 1

	fs := make([]Feature, len(vs))
	for i, v := range vs {
		h.Reset()
		h.Write(v)
		sum := h.Sum64()

		fs[i] = Feature{
			Hash:   sum,
			Weight: weight,
		}
	}
	return fs
}
