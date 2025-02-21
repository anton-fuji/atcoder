package main

import (
	"fmt"
)

func isValidPattern(S string) bool {
	patterns := []string{"ACE", "BDF", "CEG", "DFA", "EGB", "FAC", "GBD"}
	for _, pattern := range patterns {
		if S == pattern {
			return true
		}
	}
	return false
}

func main() {
	var S string
	fmt.Scan(&S)

	if isValidPattern(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
