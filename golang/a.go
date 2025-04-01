package main

import (
	"fmt"
	"strings"
)

func main() {
	var num string
	fmt.Scan(&num)

	fmt.Println(strings.Count(num, "1"))

}
