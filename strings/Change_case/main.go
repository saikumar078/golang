// Key Points:
// Character arithmetic in Go: In Go, single quotes ('a', 'h') represent runes, which are integers corresponding to their Unicode code points. For example:

// 'a' has a code point of 97
// 'h' has a code point of 104
// Expression Breakdown:

// 'h' - 'a':

// 104 - 97 = 7
// This computes the difference between the code points of 'h' and 'a'.
// 7 + 'a':

// 'a' is 97, so:
// 7 + 97 = 104
// string(104):

// Converts the rune with the code point 104 back to a string, which is "h".
// Adding to result:

// If result is a string, this operation appends "h" to result.
// Final Simplification
// The line appends "h" to result.ssd The comment // 10 + 150 is unrelated or incorrect since it doesn’t match the actual calculation. Let me know if you need further clarification!
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
			result += string(ch - 'a' + 'A') //

		} else if ch >= 'A' && ch <= 'Z' {
			result += string(ch - 'A' + 'a')
		} else {
			result += string(ch) // Non-alphabet characters remain the same
		}
	}

	fmt.Println(result)
}
