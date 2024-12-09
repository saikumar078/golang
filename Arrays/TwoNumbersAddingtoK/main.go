package main

import "fmt"

func findTwoSum(arr []int, k int) (int, int) {

	for i := 0; i < len(arr); i++ {

		for j :=i+1; j < len(arr); j++ {

			if arr[i]+arr[j] == k {
				return arr[i], arr[j] // Return the pair
			}
		}
	}
	return 0, 0 // Return 0, 0 if no pair is found
}

func main() {
	arr := []int{3, 5, 2, 8, 4}
	k := 10
	num1, num2 := findTwoSum(arr, k)
	if num1 == 0 && num2 == 0 {
		fmt.Println("No pair found")
	} else {
		fmt.Printf("The pair is: %d, %d\n", num1, num2)
	}
}
