package main

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
