package main

import (
	"fmt"
	"strings"
)

func findSmallestWord(sentence string) string {
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return ""
	}

	smallestWord := words[0]
	for _, word := range words {
		if len(word) < len(smallestWord) {
			smallestWord = word
		}
	}

	return smallestWord
}

func main() {
	sentence := "The quick brown fox jumps over the lazy dog"
	fmt.Println("Sentence:", sentence)
	fmt.Println("Smallest Word:", findSmallestWord(sentence))
}
