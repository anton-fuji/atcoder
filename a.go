package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	length := len(s) + 2
	for i := 0; i < length; i++ {
		fmt.Print("+")
	}
	fmt.Println()
	fmt.Printf("+%s+\n", s)
	for i := 0; i < length; i++ {
		fmt.Print("+")
	}
	fmt.Println()
}
