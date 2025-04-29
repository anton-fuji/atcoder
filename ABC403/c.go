package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, Q int
	fmt.Fscan(in, &n, &m, &Q)

	all := make([]bool, n+1)
	perm := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		perm[i] = make(map[int]bool)
	}

	for i := 0; i < Q; i++ {
		var typ, x, y int
		fmt.Fscan(in, &typ)
		switch typ {
		case 1:
			fmt.Fscan(in, &x, &y)
			if !all[x] {
				perm[x][y] = true
			}
		case 2:
			fmt.Fscan(in, &x)
			all[x] = true
		case 3:
			fmt.Fscan(in, &x, &y)
			if all[x] || perm[x][y] {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}
