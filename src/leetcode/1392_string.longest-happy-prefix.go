package main

import (
	"fmt"
)

func longestPrefix(s string) string {
	n := len(s)
	ans := ""
	for i := 1; i < n; i++ {
		fmt.Println(s[:i])
		fmt.Println(s[n-i:])
		if s[:i] == s[n-i:] && i > len(ans) {
			ans = s[:i]
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPrefix("level")) // l
}
