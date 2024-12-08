package main

import (
	"fmt"
)

func removeDuplicates(input string) string {
	seen := make(map[rune]bool)
	result := ""

	for _, char := range input {
		if !seen[char] {

			seen[char] = true
			result += string(char)
		}
	}
	return result
}

func main() {
	input := "programming hello"
	fmt.Println("Original:", input)
	fmt.Println("Without Duplicates:", removeDuplicates(input))
}
