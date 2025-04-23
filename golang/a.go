package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scan(&n)

	switch {
	case n == "N":
		fmt.Println("S")
	case n == "S":
		fmt.Println("N")
	case n == "E":
		fmt.Println("W")
	case n == "W":
		fmt.Println("E")
	case n == "NE":
		fmt.Println("SW")
	case n == "SW":
		fmt.Println("NE")
	case n == "SE":
		fmt.Println("NW")
	case n == "NW":
		fmt.Println("SE")
	}
}
