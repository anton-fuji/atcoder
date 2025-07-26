package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)

	var flag bool
	for i := 0; i < n; i++ {
		if a[i] == x {
			flag = true
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
