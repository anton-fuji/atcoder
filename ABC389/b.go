package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var x string
	fmt.Fscan(in, &x)

	xx := new(big.Int)
	xx.SetString(x, 10)

	n := 1
	fact := big.NewInt(1)

	for {
		if fact.Cmp(xx) == 0 {
			fmt.Println(n)
			return
		}

		n++
		fact.Mul(fact, big.NewInt(int64(n)))
	}

}
