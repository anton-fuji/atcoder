package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Gift struct {
	P, A, B int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	readInt := func() int {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		return val
	}

	N := readInt()
	gifts := make([]Gift, N)
	for i := 0; i < N; i++ {
		p := readInt()
		a := readInt()
		b := readInt()
		gifts[i] = Gift{p, a, b}
	}

	fmt.Fprintln(os.Stderr, gifts)
	Q := readInt()
	query := make([]int, Q)
	for i := 0; i < Q; i++ {
		query[i] = readInt()
	}

	cache := make(map[int]int)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, x := range query {
		if result, ok := cache[x]; ok {
			fmt.Fprintln(writer, result)
			continue
		}

		tension := x
		for _, g := range gifts {
			if tension >= g.P {
				tension += g.A
			} else {
				if tension < g.B {
					tension = 0
				} else {
					tension -= g.B
				}
			}
		}
		cache[x] = tension
		fmt.Fprintln(writer, tension)
	}
}
