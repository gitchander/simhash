package main

import (
	"fmt"

	"github.com/gitchander/simhash"
)

func main() {
	testWords()
	//testShingleRunes()
	//testShingleStrings()
	//testChinese()
	//testShingling()
}

var samples = []string{
	"this is a test phrase",
	"this is a test phrass",
	"these are test phrases",
	"foo bar",
}

func testWords() {

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		ws := simhash.Words(sample)
		fs := simhash.StringSlice(ws)

		shs[i] = simhash.Simhash(fs)
		//fmt.Println(shs[i])
	}

	for i := 1; i < len(shs); i++ {
		fmt.Println(simhash.Compare(shs[0], shs[i]))
	}
}

func testShingleRunes() {

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		p := simhash.ShinglingRunes{
			Runes: []rune(sample),
			K:     2,
		}

		shs[i] = simhash.Simhash(p)
		//fmt.Println(shs[i])
	}

	for i := 1; i < len(shs); i++ {
		fmt.Println(simhash.Compare(shs[0], shs[i]))
	}
}

func testShingleStrings() {

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		words := simhash.Words(sample)
		shingles := simhash.ShingleStrings(words, 3)
		p := simhash.StringSlice(shingles)

		shs[i] = simhash.Simhash(p)
	}

	for i := 1; i < len(shs); i++ {
		fmt.Println(simhash.Compare(shs[0], shs[i]))
	}
}

func testChinese() {

	samples := []string{
		"你好 世界 呼噜",
		"你好 世界 呼噜。",
		"this is a test phrase",
		"this is a test phrass",
		"these are test phrases",
	}

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		words := simhash.Words(sample)
		shingles := simhash.ShingleStrings(words, 2)
		p := simhash.StringSlice(shingles)

		shs[i] = simhash.Simhash(p)
	}

	for i := 1; i < len(shs); i++ {
		fmt.Println(simhash.Compare(shs[0], shs[i]))
	}
}

func testShingling() {

	text := `Я помню чудное мгновенье:
Передо мной явилась ты,
Как мимолетное виденье,
Как гений чистой красоты.
`
	words := simhash.Words(text)
	shingles := simhash.ShingleStrings(words, 3)

	for _, s := range shingles {
		fmt.Println(s)
	}
}
