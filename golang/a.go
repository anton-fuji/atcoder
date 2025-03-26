package main

import (
	"fmt"
)

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	B := make([][]int, m)
	for i := range B {
		B[i] = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, l)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(C[i][j])
		}
		fmt.Println()
	}
}
