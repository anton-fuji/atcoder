package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	nums := []int{a, b, c}

	sort.Ints(nums)

	fmt.Println(nums[0], nums[1], nums[2])
}
