pckage main

import (
	"bufio"
	"fmt"
	"os"
)

func abs() {
  
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

  h := make([]int, n)	
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
	}

  
	dp := make([]int, n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		if i+1 < n {
			d := dp[i] + abs(h[i+1]-h[i])
			if d < dp[i+1] {
				dp[i+1] = d
			}
		}

		if i+2 < n {
			d := dp[i] + abs(h[i+2]-h[i])
			if d < dp[i+2] {
				dp[i+2] = d
			}
		}
	}

	fmt.Println(dp[n-1])

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}
