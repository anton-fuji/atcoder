package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i] = sc.Text()
	}

	t := make([]string, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		t[i] = sc.Text()
	}

	for a := 0; a <= n-m; a++ {
		for b := 0; b <= n-m; b++ {
			match := true
			for i := 0; i < m; i++ {
				for j := 0; j < m; j++ {
					if s[a+i][b+j] != t[i][j] {
						match = false
						break
					}
				}
				if !match {
					break
				}
			}
			if match {
				fmt.Println(a+1, b+1)
			}
		}
	}
}
