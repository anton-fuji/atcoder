package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var n int

	fmt.Fscan(os.Stdin, &s)
	fmt.Fscan(os.Stdin, &n)
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		sc.Scan()
		t := sc.Text()
		if strings.Count(t, s) > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
