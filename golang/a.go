package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s := sc.Text()

	cnt := 0
	for i := 0; i < n-1; i++ {
		if s[i] == '#' && s[i+1] == '.' && s[i+2] == '#' {
			cnt++
		}
	}
	fmt.Println(cnt)

}
