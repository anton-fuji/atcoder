// https://onlinejudge.u-aizu.ac.jp/problems/ITP1_9_A
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var w string
	fmt.Scan(&w)

	w = strings.ToLower(w)

	// 複数入力の場合、'bufio'を使う。1行ずつ読み込む。
	sc := bufio.NewScanner(os.Stdin)
	cnt := 0

	for sc.Scan() {
		line := sc.Text()
		if line == "END_OF_TEXT" {
			break
		}

		wds := strings.Fields(line)

		for _, wd := range wds {
			if strings.ToLower(wd) == w {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
