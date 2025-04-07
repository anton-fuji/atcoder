// https://atcoder.jp/contests/abc083/tasks/abc083_b
package main

import "fmt"

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := 0
	for i := 0; i < n; i++ {
		s := digitSum(i)
		if a <= s && b >= s {
			ans += i
		}
	}
	fmt.Println(ans)
}
