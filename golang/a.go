package main

import (
	"fmt"
	"math"
)

func visit(cnt, a int) int {
	if a%2 == 1 {
		return cnt
	}
	return visit(cnt+1, a/2)
}

func main() {
	var n, a, cnt int
	min := math.MaxInt64
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)

		cnt = visit(0, a)
		if cnt < min {
			min = cnt
		}
	}
	fmt.Println(min)
}
