package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	result := ""
	for _, r := range s {
		if unicode.IsUpper(r) {
			result += string(r)
		}
	}
	fmt.Println(result)
}
