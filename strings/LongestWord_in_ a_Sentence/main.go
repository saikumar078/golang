package main

import (
	"fmt"
	"strings"
)

func LongestWordinaSentence(sentence string) string {

	longestWord := ""

	words:=strings.Fields(sentence)

	for _,val:=range words{

		if len(val) > len(longestWord){

			longestWord=val

		}
	}

	return longestWord
}

func main() {

	sentence := "The quick brown fox jumps over the lazy dog"
	fmt.Println("Longest word in the sentences :",LongestWordinaSentence(sentence))

}