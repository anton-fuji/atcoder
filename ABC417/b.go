package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	A := make([]int, n)
	B := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&B[i])
	}

	for _, b := range B {
		for i, a := range A {
			if a == b {
				A = append(A[:i], A[i+1:]...)
				break
			}
		}
	}

	for i, v := range A {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
