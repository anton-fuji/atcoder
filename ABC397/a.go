package main

import "fmt"

func main() {
	var x float32
	fmt.Scan(&x)

	switch {
	case x >= 38.0:
		fmt.Println(1)
	case x >= 37.5 && x < 38.0:
		fmt.Println(2)
	case x < 37.5:
		fmt.Println(3)
	}
}
