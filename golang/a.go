package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func zhou_length(a, b, C float64) float64 {
	var c float64
	c = a*a + b*b - 2*a*b*math.Cos(C)
	return math.Sqrt(c)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b, C float64
	fmt.Fscan(in, &a, &b, &C)

	C = C * math.Pi / 180

	s := a * b * math.Sin(C) / 2
	h := 2 * s / a
	l := zhou_length(a, b, C) + a + b

	fmt.Printf("%.8f\n", s)
	fmt.Printf("%.8f\n", l)
	fmt.Printf("%.8f\n", h)
}
