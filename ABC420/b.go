package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

const (
	inf = 1 << 60
	MOD = 998244353
)

func main() {
	/*** INIT ***/
	defer wt.Flush()
	sc.Buffer([]byte{}, inf)
	sc.Split(bufio.ScanWords)

	n, m := InputInt(), InputInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = InputString()
	}
	score := make([]int, n)
	for j := 0; j < m; j++ {
		zero, one := 0, 0
		for i := 0; i < n; i++ {
			if s[i][j] == '0' {
				zero++
			} else {
				one++
			}
		}
		if zero < one {
			for i := 0; i < n; i++ {
				if s[i][j] == '0' {
					score[i]++
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if s[i][j] == '1' {
					score[i]++
				}
			}
		}
	}
	maxscore := -1
	for i := 0; i < n; i++ {
		if maxscore < score[i] {
			maxscore = score[i]
		}
	}
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if score[i] == maxscore {
			ans = append(ans, i+1)
		}
	}
	OutSlice(ans)

}

func OutSlice(v []int) {
	for i, val := range v {
		if i != len(v)-1 {
			fmt.Fprintf(wt, "%v ", val)
		} else {
			fmt.Fprintln(wt, val)
		}
	}
}

/*** I/O ***/
func InputInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func InputString() string {
	sc.Scan()
	return sc.Text()
}

func OutInt(val int) {
	fmt.Fprintln(wt, val)
}

func OutString(s string) {
	fmt.Fprintln(wt, s)
}
