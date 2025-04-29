package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	cnt := make(map[int]int)
	for _, v := range a {
		cnt[v]++
	}

	maxValue := -1
	ans := -1
	for i, v := range a {
		if cnt[v] == 1 {
			if v > maxValue {
				maxValue = v
				ans = i + 1
			}
		}
	}
	if ans == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
