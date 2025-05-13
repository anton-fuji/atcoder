package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)

	num1, _ := strconv.Atoi(string(s[0]))
	num2, _ := strconv.Atoi(string(s[2]))

	fmt.Println(num1 * num2)

}
