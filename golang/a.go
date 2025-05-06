package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func distance(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var x1, y1, x2, y2 float64
	fmt.Fscan(in, &x1, &y1, &x2, &y2)

	d := distance(x1, y1, x2, y2)
	fmt.Printf("%.8f\n", d)
}
