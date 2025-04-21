// https://atcoder.jp/contests/typical90/tasks/typical90_d
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()
	part := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(part[0])
	w, _ := strconv.Atoi(part[1])

	A := make([][]int, h)
	rowSum := make([]int, h)
	colSum := make([]int, w)

	for i := 0; i < h; i++ {
		sc.Scan()
		line = sc.Text()
		part = strings.Split(line, " ")
		A[i] = make([]int, w)
		for j := 0; j < w; j++ {
			val, _ := strconv.Atoi(part[j])
			A[i][j] = val
			rowSum[i] += val
			colSum[j] += val
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			//行iの合計 + 列jの合計 - A[i][j]
			sum := rowSum[i] + colSum[j] - A[i][j]
			writer.WriteString(strconv.Itoa(sum))
			if j != w-1 {
				writer.WriteByte(' ')
			}
		}
		writer.WriteByte('\n')
	}

}
