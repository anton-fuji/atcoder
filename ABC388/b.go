package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, d int
	fmt.Fscan(in, &n, &d)

	t := make([]int, n)
	l := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i], &l[i])
	}

	for k := 1; k <= d; k++ {
		maxWgt := 0
		for i := 0; i < n; i++ {
			wgt := t[i] * (l[i] + k)
			if wgt > maxWgt {
				maxWgt = wgt
			}
		}
		fmt.Println(maxWgt)
	}
}
