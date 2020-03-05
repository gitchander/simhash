package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gitchander/simhash"
)

func main() {
	testFrase()
}

func testWords() {

	samples := []string{
		"this is a test phrase",
		"this is a test phrass",
		"these are test phrases",
		"foo bar",
	}

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		p := strings.Split(sample, " ")
		x := simhash.Strings(p)

		shs[i] = simhash.Simhash(x)
		//fmt.Println(shs[i])
	}

	fmt.Println(simhash.Compare(shs[0], shs[1]))
	fmt.Println(simhash.Compare(shs[0], shs[2]))
	fmt.Println(simhash.Compare(shs[0], shs[3]))
	//fmt.Println(Compare(shs[0], shs[4]))
}

func testFrase() {
	samples := []string{
		"this is a test phrase",
		"this is a test phrass",
		"these are test phrases",
		"foo bar",
		"Аргентина манит негра",
	}

	shs := make([]uint64, len(samples))

	for i, sample := range samples {

		p := simhash.ByFrase{
			Runes: []rune(sample),
			N:     2,
		}

		shs[i] = simhash.Simhash(p)
		//fmt.Println(shs[i])
	}

	fmt.Println(simhash.Compare(shs[0], shs[1]))
	fmt.Println(simhash.Compare(shs[0], shs[2]))
	fmt.Println(simhash.Compare(shs[0], shs[3]))
	fmt.Println(simhash.Compare(shs[0], shs[4]))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printJSON(v interface{}) {
	data, err := json.Marshal(v)
	checkError(err)
	fmt.Println(string(data))
}
