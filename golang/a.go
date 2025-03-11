package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	array := []int{a, b, c}
	bubbleSort(array)

	fmt.Println(array[0], array[1], array[2])
}
