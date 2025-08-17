package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	switch s {
	case "red":
		fmt.Println("SSS")
	case "blue":
		fmt.Println("FFF")
	case "green":
		fmt.Println("MMM")
	default:
		fmt.Println("Unknown")
	}
}
