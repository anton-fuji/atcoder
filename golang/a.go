package main

import (
	"fmt"
)

func main() {
	var x int

	// 無限ループ
	for i := 1; ; i++ {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", i, x)
	}
}
