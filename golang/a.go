// https://onlinejudge.u-aizu.ac.jp/problems/ITP1_9_C
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	sT := 0
	sH := 0

	// 辞書順で大きいものを比較する時は、ただの文字列比較で十分
	for i := 0; i < n; i++ {
		var cardT, cardH string
		fmt.Fscan(in, &cardT, &cardH)

		switch {
		case cardT > cardH:
			sT += 3
		case cardT < cardH:
			sH += 3
		case cardT == cardH:
			sT++
			sH++
		}
	}
	fmt.Println(sT, sH)
}
