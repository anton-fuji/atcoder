package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	if n <= 1 {
		return
	}

	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d is a prime number", n)
	} else {
		fmt.Printf("%d is not a prime number", n)
	}

}

