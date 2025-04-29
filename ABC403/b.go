package main

import (
	"fmt"
)

func main() {
	var t, u string
	fmt.Scan(&t)
	fmt.Scan(&u)

	m, l := len(t), len(u)
	for i := 0; i <= m-l; i++ {
		ok := true
		for j := 0; j < l; j++ {
			if t[i+j] != '?' && t[i+j] != u[j] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
