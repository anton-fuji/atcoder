package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)

	set := map[int64]bool{}
	for b := int64(1); b*b <= n; b++ {
		x := b * b * 2

		for x <= n {
			set[x] = true
			if x > 1e18/2 {
				break
			}
			x *= 2
		}
	}

	fmt.Println(len(set))
}
