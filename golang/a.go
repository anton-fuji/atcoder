package main

import (
	"fmt"
)

func sorting(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	num := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	sorting(num)

	for _, v := range num {
		fmt.Print(v, " ")
	}
}
