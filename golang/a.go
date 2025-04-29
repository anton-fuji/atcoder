package main

import (
	"fmt"
)

func main() {
	a := make([]int, 7)
	for i := 0; i < 7; i++ {
		fmt.Scan(&a[i])
	}

	cnt := make(map[int]int)
	for _, v := range a {
		cnt[v]++
	}

	x := 0
	y := 0
	for _, v := range cnt {
		if v >= 3 {
			x++
		}
		if v >= 2 {
			y++
		}
	}

	if x >= 1 && y >= 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
