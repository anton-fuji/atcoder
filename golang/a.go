package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotate90(grid []string) []string {
	n := len(grid)
	rotated := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			// 90度回転の変換式：新しい[i][j] = 元の[n-j-1][i]
			row[j] = grid[n-j-1][i]
		}
		rotated[i] = string(row)
	}
	return rotated
}

// gridの差分をとる
func difference(a, b []string) int {
	n := len(a)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				count++ // 色が違えば1回の変更が必要 => 必要変更回数をカウント
			}
		}
	}
	return count
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var n int
	fmt.Sscanf(sc.Text(), "%d", &n)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i] = sc.Text()
	}

	t := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		t[i] = sc.Text()
	}

	ops := n * n * 10 //十分大きい初期値
	for rot := 0; rot < 4; rot++ {
		diff := difference(s, t) // 色変更回数
		totalOps := diff + rot   // 総操作回数　＝　色変更＋回転回数
		if totalOps < ops {
			ops = totalOps
		}
		s = rotate90(s) // 次の回転状態
	}
	fmt.Println(ops)
}
