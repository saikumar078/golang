package main

import (
	"fmt"
	"strings"
	"unicode"
)

// countVowelsAndConsonants takes an input string and counts the number of vowels and consonants in it.
func countVowelsAndConsonants(input string) (int, int) {

	// Define a string containing all vowels (both uppercase and lowercase).
	vowels := "AEIOUaieou"

	// Initialize counters for vowels and consonants, and strings to store the vowels and consonants found.
	vowelsCount, consonantCount := 0, 0
	v, c := "", ""

	// Iterate over each character in the input string.
	for _, char := range input {

		// Check if the character is a letter (ignores non-letter characters like punctuation).
		if unicode.IsLetter(char) {

			// If the character is a vowel (found in the vowels string), increment the vowel count and add the character to the vowel string.
			if strings.ContainsRune(vowels, char) {
				v += string(rune(char))   // Append the vowel to the vowel string.
				vowelsCount++             // Increment the vowel counter.
			} else {
				// If the character is not a vowel, it is considered a consonant.
				c += string(rune(char))   // Append the consonant to the consonant string.
				consonantCount++           // Increment the consonant counter.
			}
		}
	}

	// Print the vowels and consonants found.
	fmt.Println("Vowels    :", v)
	fmt.Println("Consonant :", c)

	// Return the total counts of vowels and consonants.
	return vowelsCount, consonantCount
}

func main() {
	// Call the function with the input string "ajhdsj:" and store the counts in variables.
	vowelsCount, consonantCount := countVowelsAndConsonants("ajhdsj:")

	// Print the result: number of vowels and consonants.
	fmt.Println(vowelsCount, consonantCount)
}
