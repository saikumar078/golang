package main

import (
	"fmt"
	"strconv"
)

func compressString(input string) string {
	result := ""
	count := 1
	//index--- 0,1,2,3,4,5   ------ input[1]==input[0]
	//          input[2]==input[1]   
	//          input[3]==input[2]   b==a not 

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
