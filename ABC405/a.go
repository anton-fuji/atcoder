package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var r, x int
	fmt.Fscan(in, &r, &x)

	switch x {
	case 1:
		if r >= 1600 && r <= 2999 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	case 2:
		if r >= 1200 && r <= 2399 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
