package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "selamat malam"

	split(text)
	mapping(text)
}

func split(text string) {
	text = strings.ToLower(text) //make text lower case (just for learning)

	for _, input := range text {
		fmt.Printf("%c\n", input)
	}

	return
}

func mapping(text string) {
	// ip := strings.Fields(text)
	inputCount := map[string]int{} //make(map[string]int)

	for _, t := range text {
		inputCount[string(t)] += 1
	}

	fmt.Println(inputCount)
	return
}

/*
https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#character
https://www.geeksforgeeks.org/counting-number-of-repeating-words-in-a-golang-string/
*/
