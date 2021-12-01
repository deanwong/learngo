package main

import "fmt"

// O(N)
func reverseWords(s string) string {
	ans := make([]byte, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || s[i+1] == ' ' {
			for end := i; end >= start; end-- {
				ans = append(ans, s[end])
			}
			if i < len(s)-1 && s[i+1] == ' ' {
				ans = append(ans, s[i+1])
				start = i + 2
			}
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("God Ding"))
}
