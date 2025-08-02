package main

import "fmt"

func main() {
	var n, a, b int
	var s string
	fmt.Scan(&n, &a, &b)
	fmt.Scan(&s)

	aa := a
	bb := len(s) - b
	new_s := s[aa:bb]
	fmt.Println(new_s)
}
