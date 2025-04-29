package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1 X Y: ユーザX にコンテストページ Y の閲覧権限を付与する
// 2 X: ユーザX にすべてのコンテストページの閲覧権限を付与する
// 3 X Y: ユーザX がコンテストページYを閲覧できるかを答える

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	// main 関数の最後に必ずoutのバッファを捨てずに書き出す（Flush）ように予約しておく
	defer out.Flush()

	var n, m, Q int
	fmt.Fscan(in, &n, &m, &Q)

	//全権限のフラグ
	all := make([]bool, n+1)

	// 個別権限付与マップの初期化
	perm := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		perm[i] = make(map[int]bool)
	}

	//各クエリ実行
	for i := 0; i < Q; i++ {
		var typ, x, y int
		fmt.Fscan(in, &typ)
		switch typ {
		case 1:
			fmt.Fscan(in, &x, &y)
			if !all[x] {
				perm[x][y] = true
			}
		case 2:
			fmt.Fscan(in, &x)
			all[x] = true
		case 3:
			fmt.Fscan(in, &x, &y)
			if all[x] || perm[x][y] {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}
