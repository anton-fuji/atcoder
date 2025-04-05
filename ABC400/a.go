package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	row := 400 / a

	if 400%a == 0 {
		fmt.Println(row)
	} else {
		fmt.Println(-1)
	}
}
