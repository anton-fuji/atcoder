package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	var n int
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var q string
		var a, b int
		fmt.Fscan(in, &q, &a, &b)

		switch q {
		case "print":
			fmt.Println(s[a : b+1])
			// 一文字ずつ扱うときは[]rune
		case "reverse":
			sub := []rune(s[a : b+1])
			for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
				sub[i], sub[j] = sub[j], sub[i]
			}
			s = s[:a] + string(sub) + s[b+1:]
		case "replace":
			var p string
			fmt.Fscan(in, &p)
			s = s[:a] + p + s[b+1:]
		}
	}
}
