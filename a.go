package main

import (
	"fmt"
)

func countStrings(s []string) []string {
	count := map[string]bool{}
	result := []string{}

	for _, v := range s {
		if !count[v] {
			count[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	s := []string{"apple", "grape", "banana", "apple"}

	fmt.Println(countStrings(s))
}
