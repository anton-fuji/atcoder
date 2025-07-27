package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(input, &s)

	n := len(s)
	t := make([]rune, n)

	// t に対して#, .で初期化
	for i, v := range s {
		if v == '#' {
			t[i] = '#'
		} else {
			t[i] = '.'
		}
	}

	// 最後にoを配置したインデックスを記録
	lastIdx := -1
	// 最後に#が見つかったインデックスを記録
	lastHash := -1

	for i := 0; i < n; i++ {
		if s[i] == '#' {
			t[i] = '#'
			lastHash = i
		} else {
			if lastIdx == -1 || lastHash > lastIdx {
				t[i] = 'o'
				lastIdx = i
			} else {
				t[i] = '.'
			}
		}
	}

	fmt.Println(string(t))
}
