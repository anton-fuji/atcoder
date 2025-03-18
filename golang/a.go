package main

import (
	"fmt"
)

func main() {
	ranks := []string{"S", "H", "C", "D"}
	cardExists := make(map[string]bool)

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var rank string
		var num int
		fmt.Scan(&rank, &num)
		cardExists[fmt.Sprintf("%s %d", rank, num)] = true
	}

	for _, rank := range ranks {
		for num := 1; num <= 13; num++ {
			card := fmt.Sprintf("%s %d", rank, num)
			if !cardExists[card] {
				fmt.Println(card)
			}
		}
	}
}
