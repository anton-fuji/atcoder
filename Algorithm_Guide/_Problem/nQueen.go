package main

import "fmt"

// ここにQが置けるか判定
func isSafe(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]-i == col-row || board[i]+i == col+row {
			return false
		}
	}
	return true
}

// Qの配置
func solveQueens(n, row int, board []int, solutions *[][]int) {
	if row == n {
		solution := make([]int, n)
		copy(solution, board)
		*solutions = append(*solutions, solution)
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solveQueens(n, row+1, board, solutions)
			board[row] = -1
		}
	}
}

// Nクイーンを実行
func nQueens(n int) [][]int {
	board := make([]int, n)
	for i := range board {
		board[i] = -1
	}
	solutions := [][]int{}
	solveQueens(n, 0, board, &solutions)
	return solutions
}

// 結果
func printSolutions(solutions [][]int, n int) {
	for _, sol := range solutions {
		fmt.Println("Ans.")
		for i := 0; i < n; i++ {
			row := ""
			for j := 0; j < n; j++ {
				if sol[i] == j {
					row += "Q"
				} else {
					row += "."
				}
			}
			fmt.Println(row)
		}
		fmt.Println()
	}
}

func main() {
	n := 4
	solutions := nQueens(n)
	fmt.Printf("N=%d のAnsの数: %d\n", n, len(solutions))
	printSolutions(solutions, n)
}
