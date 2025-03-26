package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&B[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(A[i][j] + B[i][j])
		}
		fmt.Println()
	}
}
