package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	counts := make([]int, 26)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			if unicode.IsLetter(c) {
				lower := unicode.ToLower(c)
				index := lower - 'a'
				counts[index]++
			}
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", 'a'+i, counts[i])
	}
}
