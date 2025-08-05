package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	T := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	if T[0] > s {
		fmt.Println("No")
		return
	}

	for i := 0; i < n-1; i++ {
		if T[i+1]-T[i] > s {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
