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

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// xを大きい方から順に試す
	for x := n; x >= 0; x-- {
		cnt := 0
		for _, v := range a {
			// x以上の要素をカウント
			if v >= x {
				cnt++
			}
		}
		if cnt >= x {
			fmt.Fprintln(out, x)
			return
		}
	}
}
