package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	var n, l, r int
	var s string
	fmt.Fscan(input, &n, &l, &r)
	fmt.Fscan(input, &s)

	judge := true
	for i := l - 1; i < r; i++ {
		if s[i] != 'o' {
			judge = false
			break
		}
	}

	if judge {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
