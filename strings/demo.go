package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	// Convert to lowercase and sort both strings
	s1 := strings.Split(strings.ToLower(str1), "")
	s2 := strings.Split(strings.ToLower(str2), "")
	sort.Strings(s1)
	sort.Strings(s2)

	return strings.Join(s1, "") == strings.Join(s2, "")
}

func main() {
	
	s2 :=strings.Split("the","")
	s3 :=strings.Split("het","")

	fmt.Println(s2)
	fmt.Println(s3)

	sort.Strings(s2)
	sort.Strings(s3)
	fmt.Println(s2)
	fmt.Println(s3)

	strings.Compare(s2,s3)
}
