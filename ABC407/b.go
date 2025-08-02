package main

import (
	"fmt"
)

func main() {
	var x, y, total int
	fmt.Scan(&x, &y)

	for a := 1; a <= 6; a++ {
		for b := 1; b <= 6; b++ {
			sum := a + b
			diff := abs(a - b)

			if sum >= x || diff >= y {
				total++
			}
		}
	}
	fmt.Printf("%.15f\n", float64(total)/36.0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
