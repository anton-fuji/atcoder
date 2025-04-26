package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = '-'
	}
	if n%2 == 0 {
		res[n/2-1] = '='
		res[n/2] = '='
	} else {
		res[n/2] = '='
	}
	fmt.Println(string(res))
}
