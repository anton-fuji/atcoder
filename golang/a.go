package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	for _, c := range input {
		if unicode.IsUpper(c) {
			fmt.Printf("%c", unicode.ToLower(c))
		} else {
			fmt.Printf("%c", unicode.ToUpper(c))
		}
	}
}
