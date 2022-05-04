package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	lookup := make(map[byte]int)
	i := 0
	for j := 0; j < n; j++ {
		lookup[s[j]]++
		for i <= j && lookup[s[j]] > 1 {
			lookup[s[i]]--
			i++
		}
		temp := j - i + 1
		if temp > ans {
			ans = temp
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}
