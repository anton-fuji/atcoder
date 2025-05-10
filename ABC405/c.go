package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	totalSum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		totalSum += a[i]
	}

	sum := 0
	for i := 0; i < n; i++ {
		totalSum -= a[i]
		sum += a[i] * totalSum
	}

	fmt.Println(sum)
}
