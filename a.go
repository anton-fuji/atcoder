package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	for i := 0; i < n-1; i++ {
		if m >= a[i] {
			sum += a[i]
		}
	}
	fmt.Println(sum)
}
