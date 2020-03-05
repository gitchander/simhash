package main

import (
	"fmt"

	"github.com/gitchander/simhash"
)

func main() {
	testWords()
	//testRunesGroup()
	//testStringsGroup()
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

	fmt.Println(simhash.Compare(shs[0], shs[1]))
	fmt.Println(simhash.Compare(shs[0], shs[2]))
	fmt.Println(simhash.Compare(shs[0], shs[3]))
}

func testRunesGroup() {

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		p := simhash.RunesGroup{
			Runes: []rune(sample),
			N:     3,
		}

		shs[i] = simhash.Simhash(p)
		//fmt.Println(shs[i])
	}

	fmt.Println(simhash.Compare(shs[0], shs[1]))
	fmt.Println(simhash.Compare(shs[0], shs[2]))
	fmt.Println(simhash.Compare(shs[0], shs[3]))
}

func testStringsGroup() {

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		p := simhash.StringsGroup{
			Strings: simhash.Words(sample),
			N:       3,
		}

		shs[i] = simhash.Simhash(p)
	}

	fmt.Println(simhash.Compare(shs[0], shs[1]))
	fmt.Println(simhash.Compare(shs[0], shs[2]))
	fmt.Println(simhash.Compare(shs[0], shs[3]))
}
