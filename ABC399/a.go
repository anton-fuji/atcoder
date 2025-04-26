package main

import (
	"fmt"
)

func main() {
	var n int
	var s, t string
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Scan(&t)

	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
