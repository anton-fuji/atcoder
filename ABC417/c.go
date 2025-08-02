package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	cntMap := make(map[int]int)
	for i := 0; i < n; i++ {
		key := i + A[i]
		cntMap[key]++
	}

	var ans int64
	for j := 0; j < n; j++ {
		key := j - A[j]
		ans += int64(cntMap[key])
	}

	fmt.Println(ans)
}
