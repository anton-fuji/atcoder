package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	var peak []int
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			curr := matrix[i][j]
			isPeak := true
			for d := 0; d < 4; d++ {
				ni := i + dy[d]
				nj := j + dx[d]
				if ni >= 0 && ni < n && nj >= 0 && nj < n {
					if matrix[ni][nj] >= curr {
						isPeak = false
						break
					}
				}
			}
			if isPeak {
				peak = append(peak, curr)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(peak)))

	for _, h := range peak {
		fmt.Println(h)
	}
}
