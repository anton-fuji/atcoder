package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	cnt := make(map[int]int)
	for i := 0; i < 4; i++ {
		var a int
		fmt.Fscan(in, &a)
		cnt[a]++
	}

	has1, has2, has3 := 0, 0, 0
	for _, v := range cnt {
		switch v {
		case 1:
			has1++
		case 2:
			has2++
		case 3:
			has3++
		}
	}

	if (has3 == 1 && has1 == 1) || has2 == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
