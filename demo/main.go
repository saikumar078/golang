package main

import (
	"fmt"
)

func findDuplicateCharacters(s string) {
	counts := make(map[rune]int)

	// Count occurrences of each character
	for _, char := range s {
		counts[char]++
	}

	// Print duplicates
	fmt.Println("Duplicate characters:")
	for char, count := range counts {
		if count > 1 {
			fmt.Printf("'%c' occurs %d times\n", char, count)
		}
	}
}

func main() {
	input := "programming"
	findDuplicateCharacters(input)
}
