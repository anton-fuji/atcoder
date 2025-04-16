package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	A := make([]int, n+1)

	for i := 0; i < k && i =< n; i++ {
			A[i] = 1
	}

	sum := 0
	for i := k; i <= n; i++ {
		sum += A[i]
	}
	fmt.Println()

}
