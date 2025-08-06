package main

import (
	"fmt"
	"sort"
)

func removeDup(n []int) []int {
	judge := map[int]bool{}
	result := []int{}

	for _, v := range n {
		_, ok := judge[v]
		if !ok {
			judge[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	dup := removeDup(A)
	sort.Slice(dup, func(i, j int) bool {
		return dup[i] < dup[j]
	})

	fmt.Println(len(dup))
	for i, v := range dup {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
