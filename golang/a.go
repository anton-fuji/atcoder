package main

import (
	"fmt"
)

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	// 配列の初期化
	table := make([][]int, r)
	for i := 0; i < r; i++ {
		table[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&table[i][j])
		}
	}

	// 列ごとの合計用
	colSums := make([]int, c)
	totalSum := 0

	// 行ごとの出力と合計
	for i := 0; i < r; i++ {
		rowSum := 0
		for j := 0; j < c; j++ {
			val := table[i][j]
			fmt.Printf("%d ", val)
			rowSum += val
			colSums[j] += val
			totalSum += val
		}
		fmt.Println(rowSum)
	}

	// 最後の行（列ごとの合計 + 全体合計）
	for j := 0; j < c; j++ {
		fmt.Printf("%d ", colSums[j])
	}
	fmt.Println(totalSum)
}
