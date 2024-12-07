package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountVowelsandConsonants(s string) {

	vowels := "AEIOUaeiou"
	 vcount := 0
	 c_count := 0 
	 vowels_letter:=""
	 consonants_letter:=""
	 
	 
	for _, char := range s {
		if unicode.IsLetter(char) {

			if strings.Contains(vowels, string(char)) {
                 vowels_letter+=string(char)
				 
				vcount++
			} else {
				c_count++
				consonants_letter+=string(char)
				 
			}
		}
	}
	fmt.Println("vowels_count:", vcount, "consonantcount:", c_count)

	fmt.Println("vowels_letter: ",vowels_letter)    
	fmt.Println("consonants_letter: ",consonants_letter)    

}

func main() {

	CountVowelsandConsonants("helloWorld!!")
}
