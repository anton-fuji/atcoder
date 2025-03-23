package main

import "fmt"

func main() {
	for {
		var m, f, r int
		fmt.Scan(&m, &f, &r)

		if m == -1 && f == -1 && r == -1 {
			break
		}

		switch {
		case m == -1 || f == -1:
			fmt.Println("F")
		case m+f >= 80:
			fmt.Println("A")
		case 65 <= m+f:
			fmt.Println("B")
		case 50 <= m+f:
			fmt.Println("C")
		case 30 <= m+f:
			if r >= 50 {
				fmt.Println("C")
			} else {
				fmt.Println("D")
			}
		default:
			fmt.Println("F")
		}
	}
}
