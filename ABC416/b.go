package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	t := make([]rune, n)

	for i, v := range s {
		if v == '#' {
			t[i] = '#'
		} else {
			t[i] = '.'
		}
	}

	lastOIdx := -1
	lastHashIdx := -1

	for i := 0; i < n; i++ {
		if s[i] == '#' {
			t[i] = '#'
			lastHashIdx = i
		} else {
			if lastOIdx == -1 || lastHashIdx > lastOIdx {
				t[i] = 'o'
				lastOIdx = i
			} else {
				t[i] = '.'
			}
		}
	}

	fmt.Println(string(t))
}
