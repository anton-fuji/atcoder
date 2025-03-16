package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var H, W int
		fmt.Scan(&H, &W)

		if H == 0 && W == 0 {
			break
		}

		for i := 0; i < H; i++ {
			fmt.Println(strings.Repeat("#", W))
		}
		fmt.Println()
	}
}
