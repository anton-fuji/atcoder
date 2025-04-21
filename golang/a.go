// https://onlinejudge.u-aizu.ac.jp/problems/ITP1_9_B
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		sc.Scan()
		s := sc.Text()
		if s == "-" {
			break
		}

		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())

		for i := 0; i < m; i++ {
			sc.Scan()
			h, _ := strconv.Atoi(sc.Text())

			// シャッフル処理: 左から h 文字を末尾に移動
			s = s[h:] + s[:h]
		}
		fmt.Println(s)
	}

}
