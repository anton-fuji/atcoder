package main

import "fmt"

func main() {
	var n [7]int

	for i := 0; i < 7; i++ {
		fmt.Scan(&n[i])
	}

	freq := make(map[int]int)
	for _, v := range n {
		freq[v]++
	}

	x, y := 0, 0
	for _, cnt := range freq {
		if cnt >= 3 {
			x++
		}
		if cnt >= 2 {
			y++
		}
	}

	if x >= 1 && y >= 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
