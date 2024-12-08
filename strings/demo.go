package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(input string) {

	words := strings.Fields(input)
	longword := ""

	for _, word := range words {

		if len(word) > len(longword) {
			longword=word
		}
	
	}
	fmt.Println(longword)

}

func main() {
	removeDuplicates(" kjxdjb kjvdbnx sai,q progeaminf")
}
