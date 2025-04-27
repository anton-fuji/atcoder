package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	ans := 0
	pos := 0

	for i := 0; pos < n; i++ {
		e := 'i'
		if i%2 == 1 {
			e = 'o'
		}
		if rune(s[pos]) == e {
			pos++
		}
		ans++
	}
	if ans%2 == 1 {
		ans++
	}
	// 追加すべき個数＝理想形の文字列(ans) - 元々の文字列の長さ
	fmt.Println(ans - n)

}
