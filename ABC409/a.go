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
	var t, a string

	fmt.Fscan(in, &n, &t, &a)

	var flag bool
	for i := 0; i < n; i++ {
		if t[i] == 'o' && a[i] == 'o' {
			flag = true
			break
		} else {
			flag = false
		}
	}

	if flag {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}

}
