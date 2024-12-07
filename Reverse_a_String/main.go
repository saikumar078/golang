package main

import "fmt"

func isRevers(s string) {

	Result := ""

	// for i:=len(s)-1;i>=0;i--{
	// 	Result+=string(s[i])
	// }

	for _, val := range s {
		Result = string(val) + Result
	}

	fmt.Println(Result)
}

func main() {
	isRevers("hello")
}
