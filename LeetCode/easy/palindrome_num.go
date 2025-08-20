// https://leetcode.com/problems/palindrome-number/description/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := isPalindrome(n)
	fmt.Println(res)
}

// without pattern of int → string
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revNum := 0

	for x > revNum {
		revNum = revNum*10 + x%10
		x /= 10
	}

	return x == revNum || x == revNum/10
}
