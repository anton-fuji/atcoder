package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	queue := []int{}
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < Q; i++ {
		sc.Scan()
		line := sc.Text()
		parts := strings.Fields(line)

		if parts[0] == "1" {
			x, _ := strconv.Atoi(parts[1])
			queue = append(queue, x)
		} else if parts[0] == "2" {
			fmt.Println(queue[0])
			queue = queue[1:]
		}
	}
}
