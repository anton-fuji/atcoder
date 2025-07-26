package main

import (
	"fmt"
	"math"
	"sort"
)

var all []string
var s []string
var N_global, K_global int

func combinations(currS string, currD int) {
	if currD == K_global {
		all = append(all, currS)
		return
	}

	for i := 0; i < N_global; i++ {
		combinations(currS+s[i], currD+1)
	}
}

func main() {
	var N, K, x int
	fmt.Scan(&N, &K, &x)

	N_global = N
	K_global = K

	s = make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s[i])
	}

	maxCombinations := int(math.Pow(float64(N), float64(K)))
	all = make([]string, 0, maxCombinations)

	combinations("", 0)
	sort.Strings(all)

	fmt.Println(all[x-1])

}
