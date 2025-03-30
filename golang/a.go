package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "0" {
			break
		}

		sum := 0
		for _, c := range line {
			sum += int(c - '0')
		}

		fmt.Println(sum)
	}
}
