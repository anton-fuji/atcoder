package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	minNum := math.MaxInt64
	maxNum := math.MinInt64
	sum := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
		sum += num
	}
	fmt.Printf("%d %d %d\n", minNum, maxNum, sum)
}
