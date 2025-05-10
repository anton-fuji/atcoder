// package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var pr [26]bool
	for _, v := range s {
		pr[v-'a'] = true
	}

	for i := 0; i < 26; i++ {
		if !pr[i] {
			fmt.Printf("%c\n", 'a'+i)
			return
		}
	}
}


{*<解答>*}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Scan()
// 	s := sc.Text()

// 	seen := make(map[rune]bool)
// 	for _, c := range s {
// 		seen[c] = true
// 	}

// 	for c := 'a'; c <= 'z'; c++ {
// 		if !seen[c] {
// 			fmt.Println(string(c))
// 			return
// 		}
// 	}
// }
