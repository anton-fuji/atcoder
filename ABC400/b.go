package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var limit = int64(1e9)

	sum := int64(1)
	ans := int64(1)

	for i := 0; i < m; i++ {
		if ans > limit/int64(n) {
			fmt.Println("inf")
			return
		}
		ans *= int64(n)

		if sum > limit-ans {
			fmt.Println("inf")
			return
		}
		sum += ans
	}
	fmt.Println(sum)
}
