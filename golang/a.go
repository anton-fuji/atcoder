package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	maxNum := math.MinInt64
	minNum := math.MaxInt64
	sum := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
		sum += num
	}
	fmt.Println(minNum, maxNum, sum)
}
