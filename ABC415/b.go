package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	pos := []int{}
	for i, c := range s {
		if c == '#' {
			pos = append(pos, 1+i)
		}
	}

	for i := 0; i < len(pos); i += 2 {
		fmt.Printf("%d,%d\n", pos[i], pos[i+1])
	}
}
