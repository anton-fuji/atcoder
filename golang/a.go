package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				total := i*500 + j*100 + k*50
				if total == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
