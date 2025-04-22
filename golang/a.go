package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sc.Scan()
	n := sc.Text()

	cnt := 0
	for _, s := range n {
		if s == '.' {
			cnt++
		}
	}
	fmt.Println(cnt)

}
