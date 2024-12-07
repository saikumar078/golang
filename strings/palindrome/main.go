package main

import "fmt"

func pallidrome(s string) {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}

	if result == s {
		fmt.Println("is pallindrome")
	}else{
		fmt.Println("not Palindrome!!")
	}
}
func main() {
	pallidrome("madam")
}
