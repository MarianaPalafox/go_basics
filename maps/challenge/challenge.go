package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := strings.Fields(text)

	// My solution
	wordFreq := map[string]int{}

	for _, word := range words {
		freq, ok := wordFreq[strings.ToLower(word)]
		if !ok {
			wordFreq[strings.ToLower(word)] = 1
		} else {
			wordFreq[strings.ToLower(word)] = freq + 1
		}
	}

	fmt.Println(wordFreq)

	fmt.Println("__________________")

	// Miki's solution
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++ // because when the key value pair doen't exists it returns 0 duuh
	}
	fmt.Println(counts)
}
