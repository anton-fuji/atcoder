package main

import (
	"fmt"
)

func judge(a, b string) int {
	if a == b {
		return 0
	}

	if (a == "G" && b == "C") ||
		(a == "C" && b == "P") ||
		(a == "P" && b == "G") {
		return 1
	}
	return 2
}

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		var a, b string
		fmt.Scan(&a, &b)

		result := judge(a, b)
		if result == 1 {
			count += 1
		}
	}
	fmt.Println(count)
}
