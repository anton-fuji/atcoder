package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	var n int
	var s, t string
	var cnt int

	fmt.Fscan(input, &n)
	fmt.Fscan(input, &s, &t)

	runes_S := []rune(s)
	runes_T := []rune(t)
	for i := 0; i < n; i++ {
		if runes_S[i] != runes_T[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
