package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	max := 0.0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subStr := s[i : j+1]

			if len(subStr) >= 3 {
				if subStr[0] == 't' && subStr[len(subStr)-1] == 't' {
					tCnt := strings.Count(subStr, "t")
					fillRate := float64(tCnt-2) / float64(len(subStr)-2)
					if fillRate > max {
						max = fillRate
					}
				}
			}
		}
	}
	fmt.Printf("%.10f\n", max)
}
