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

	var x, y int
	fmt.Fscan(in, &x, &y)

	tuki := x + y
	if tuki > 12 {
		ans := tuki - 12
		fmt.Fprintln(out, ans)
	} else {
		fmt.Fprintln(out, tuki)
	}
}
