package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// bufio, os使用パターン
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.ParseInt(sc.Text(), 10, 64)
	sc.Scan()
	m, _ := strconv.ParseInt(sc.Text(), 10, 64)

	x, pow := int64(0), int64(1)
	limit := int64(1e9)

	for i := int64(0); i <= m; i++ {
		if x > limit-pow {
			fmt.Println("inf")
			return
		}
		x += pow
		if i < m && pow < limit/n {
			fmt.Println("inf")
			return
		}
		pow *= n
	}
	fmt.Println(x)
}
