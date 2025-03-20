package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// create matrix_A
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	// creat fact_B
	B := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&B[i])
	}

	// c = A[][] * B[]
	c := make([]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += A[i][j] * B[j]
		}
		c[i] = sum
	}

	for i := 0; i < n; i++ {
		fmt.Println(c[i])
	}
}
