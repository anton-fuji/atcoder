package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)

	result := make([]int, 0, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var x int
			fmt.Fscan(in, &x)
			result = append(result, x)
		} else {
			sort.Slice(result, func(i, j int) bool {
				return result[i] < result[j]
			})
			fmt.Fprintln(out, result[0])
			result = result[1:]
		}
	}
}
