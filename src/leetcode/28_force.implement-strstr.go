package main

import "fmt"

// O(n)
func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	for i := 0; i <= m-n; i++ {
		// s1出发点
		a := i
		b := 0
		for b < n && haystack[a] == needle[b] {
			a++
			b++
		}
		if n == b {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))  // 2
	fmt.Println(strStr("aaaaa", "bba")) // -1
	fmt.Println(strStr("", ""))         // 0
}
