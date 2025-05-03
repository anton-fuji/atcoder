package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	loggedIn := false
	errCnt := 0

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)

		switch s {
		case "login":
			loggedIn = true
		case "logout":
			loggedIn = false
		case "public":
		case "private":
			if !loggedIn {
				errCnt++
			}
		}
	}
	fmt.Println(errCnt)
}
