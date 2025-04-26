package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	ranks := make([]int, n)
	currentRank := 1

	for {
		maxScore := -1
		for i := 0; i < n; i++ {
			if ranks[i] == 0 && scores[i] > maxScore {
				maxScore = scores[i]
			}
		}

		if maxScore == -1 {
			break
		}

		count := 0
		for i := 0; i < n; i++ {
			if scores[i] == maxScore && ranks[i] == 0 {
				ranks[i] = currentRank
				count++
			}
		}

		currentRank += count
	}

	for i := 0; i < n; i++ {
		fmt.Println(ranks[i])
	}
}
