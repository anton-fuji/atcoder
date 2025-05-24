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

	var h, w int
	fmt.Fscan(in, &h, &w)

	a := make([][]int64, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int64, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	rowSum := make([]int64, h)
	colSum := make([]int64, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			v := a[i][j]
			rowSum[i] += v
			colSum[j] += v
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans := rowSum[i] + colSum[j] - a[i][j]
			fmt.Fprint(out, ans)
			if j < w-1 {
				out.WriteByte(' ')
			}
		}
		out.WriteByte('\n')
	}

}
