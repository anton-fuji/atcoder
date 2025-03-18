package main

import "fmt"

func main() {
	appartment := make([][][]int, 4)
	for i := range appartment {
		appartment[i] = make([][]int, 3)
		for j := range appartment[i] {
			appartment[i][j] = make([]int, 10)
		}
	}

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scan(&b, &f, &r, &v)
		appartment[b-1][f-1][r-1] += v
	}

	for b := 0; b < 4; b++ {
		if b > 0 {
			fmt.Println("####################")
		}
		for f := 0; f < 3; f++ {
			for r := 0; r < 10; r++ {
				fmt.Print(" ", appartment[b][f][r])
			}
			fmt.Println()
		}

	}
}
