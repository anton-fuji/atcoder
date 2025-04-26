package main

import (
	"fmt"
)

func main() {
	var n [7]int

	for i := 0; i < len(n); i++ {
		fmt.Scan(&n[i])
	}

	freq := make(map[int]int)
	for _, v := range n {
		freq[v]++
		fmt.Println(freq)
	}

}
