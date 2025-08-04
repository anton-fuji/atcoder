package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s, result string
	fmt.Scan(&s)

	for _, u := range s {
		if unicode.IsUpper(u) {
			result += string(u)
		}
	}
	fmt.Println(result)
}
