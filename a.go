package main

import (
	"fmt"
	"strconv"
)

func checkNum(n int) bool {
	sNum := strconv.Itoa(n)

	cnt := make(map[rune]int)
	for _, digit := range sNum {
		cnt[digit]++
	}
	return cnt['1'] == 1 && cnt['2'] == 2 && cnt['3'] == 3
}

func main() {
	var n int
	fmt.Scan(&n)

	if checkNum(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
