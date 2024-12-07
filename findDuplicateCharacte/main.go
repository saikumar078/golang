package main

import "fmt"

func findDuplicateCharacte(input string) {

	maps := make(map[rune]int)

	for _, val := range input {

		maps[val]++
	}

	 for k,v := range maps{
        
		if v > 1 {
			fmt.Printf("'%c' occurs %d times\n", k,v)
		}
	 }
	
}


func main(){

	findDuplicateCharacte("hello")
}