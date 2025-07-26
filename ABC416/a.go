package main

import "fmt"

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)

	var s string
	fmt.Scan(&s)

	judge := true
	for i := l - 1; i < r; i++ {
		if s[i] != 'o' {
			judge = false
			break
		}
	}

	if judge {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
