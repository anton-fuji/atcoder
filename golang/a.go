package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	A := make([]int, n)
	count := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
		count[A[i]]++
	}

	uniqueCount := len(count)

	duplicates := n - uniqueCount

	if duplicates%2 == 1 {
		uniqueCount--
	}

	fmt.Println(uniqueCount)
}
