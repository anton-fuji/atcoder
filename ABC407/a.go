package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ratio := float64(a) / float64(b)
	ans := int(math.Round(ratio))

	fmt.Println(ans)
}
