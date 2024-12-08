package main

import (
	"fmt"
)

func main() {
	// Input string
	str := "Hello World!1"

	// Change case: Convert lowercase to uppercase and vice versa
	var result string
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {

			result += string(ch - 'a' + 'A')// 10+150

		} else if ch >= 'A' && ch <= 'Z' {
			result += string(ch - 'A' + 'a')
		} else {
			result += string(ch) // Non-alphabet characters remain the same
		}
	}

	fmt.Println(result)
}
