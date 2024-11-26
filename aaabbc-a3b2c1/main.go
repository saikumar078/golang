package main

import (
	"fmt"
	"strconv"
)

func compressString(input string) string {
	result := ""
	count := 1
	
	for i := 1; i <len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			result += string(input[i-1])+strconv.Itoa(count)
			count = 1
		}
	}
	// Append the last character and its count
	result += string(input[len(input)-1]) +strconv.Itoa(count)
	return result
}

func main() {
	input := "aaabb"
	fmt.Println(compressString(input)) // Output: a3b2
}
