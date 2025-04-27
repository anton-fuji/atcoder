package main

import (
	"fmt"
	"strings"
)

func main() {
	var t, u string
	fmt.Scan(&t)
	fmt.Scan(&u)

	var s string
	if strings.HasPrefix(t, u) {
		s = t[:len(t)-len(u)]
	}
	fmt.Println(s)
}
