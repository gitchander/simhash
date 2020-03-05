package main

import (
	"fmt"

	"github.com/gitchander/simhash"
)

func main() {
	testWords()
	//testRunesGroup()
	//testStringsGroup()
	//testChinese()
	//testGroup()
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

func testRunesGroup() {

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		p := simhash.RunesGroup{
			Runes: []rune(sample),
			N:     2,
		}

		shs[i] = simhash.Simhash(p)
		//fmt.Println(shs[i])
	}

	for i := 1; i < len(shs); i++ {
		fmt.Println(simhash.Compare(shs[0], shs[i]))
	}
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
		rs := simhash.RuneSlice([]rune(sample))
		shs[i] = simhash.Simhash(rs)
	}

	for i := 1; i < len(shs); i++ {
		fmt.Println(simhash.Compare(shs[0], shs[i]))
	}
}

func testGroup() {

	text := `Я помню чудное мгновенье:
Передо мной явилась ты,
Как мимолетное виденье,
Как гений чистой красоты.
`
	// g := simhash.RunesGroup{
	// 	Runes: []rune(text),
	// 	N:     5,
	// }

	g := simhash.StringsGroup{
		Strings: simhash.Words(text),
		N:       3,
	}

	n := g.Len()
	for i := 0; i < n; i++ {
		data := g.Bytes(i)
		fmt.Printf("%q\n", string(data))
	}
}
