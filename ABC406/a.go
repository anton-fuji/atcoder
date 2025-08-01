package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	if a == c {
		if b >= d {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else if a > c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
