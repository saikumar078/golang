package main

import "fmt"

func main() {

	a := [3]int{3, 4, 1}

	for i := 0; i < len(a); i++ {

		for j := i+1; j < len(a); j++ {

			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
		
	}
	fmt.Println(a)
	

}