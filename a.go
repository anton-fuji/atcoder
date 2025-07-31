package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var q int
	fmt.Scan(&q)

	stack := make([]int, 100)
	for i := range stack {
		stack[i] = 0
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	nexInt := func() int {
		sc.Scan()
		var x int
		fmt.Sscan(sc.Text(), &x)
		return x
	}

	for i := 0; i < q; i++ {
		cmd := nexInt()
		if cmd == 1 {
			x := nexInt()
			stack = append(stack, x)
		} else if cmd == 2 {
			top := stack[len(stack)-1]
			fmt.Println(top)
			stack = stack[:len(stack)-1]
		}
	}

}
