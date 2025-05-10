package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	cnts := make(map[int]int)
	for _, v := range a {
		if v >= 1 && v <= m {
			cnts[v]++
		}
	}

	for i := 1; i <= m; i++ {
		if cnts[i] == 0 {
			fmt.Println(0)
			return
		}
	}
	removeCnts := 0
	for i := len(a) - 1; i >= 0; i-- {
		num := a[i]
		if num >= 1 && num <= m {
			cnts[num]--
			removeCnts++
			if cnts[num] == 0 {
				break
			}
		}
	}
	fmt.Println(removeCnts)

}
