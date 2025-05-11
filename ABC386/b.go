package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	const inf = 1 << 30
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = inf
	}

	for i := 0; i < n; i++ {
		if dp[i+1] > dp[i]+1 {
			dp[i+1] = dp[i] + 1
		}

		if i+1 < n && s[i] == '0' && s[i+1] == '0' {
			if dp[i+2] > dp[i]+1 {
				dp[i+2] = dp[i] + 1
			}
		}
	}
	fmt.Println(dp[n])
}
