package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	len := len(s)
	ans := make([]byte, len)
	// rolling array
	for i := 0; i < len; i++ {
		// 转换为向右偏移 i+len-n
		ans[(i+len-n)%len] = s[i]
	}
	return string(ans)
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))    // cdefgab
	fmt.Println(reverseLeftWords("lrloseumgh", 6)) // umghlrlose
}
