package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1 X Y: ユーザX にコンテストページY の閲覧権限を付与する。
// 2 X: ユーザ X にすべてのコンテストページの閲覧権限を付与する。
// 3 X Y: ユーザ X がコンテストページY を閲覧できるかを答える。

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < q; i++ {
		if sc.Scan() {

		}
	}
}
