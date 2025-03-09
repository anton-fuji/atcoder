package main

import (
	"fmt"
)

var s string

func main() {
	fmt.Scan(&s)

	firstA := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			firstA = i
			break
		}
	}

	lastZ := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'Z' {
			lastZ = i
			break
		}
	}

	fmt.Println(lastZ - firstA + 1)
}
