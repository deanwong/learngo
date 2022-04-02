package main

import (
	"fmt"
)

// O(n)
func lengthOfLongestSubstring2(s string) int {
	length := 0
	end := -1
	dict := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		// delete element before left pointer
		if i != 0 {
			delete(dict, s[i-1])
		}
		// not found & move right pointer
		for end+1 < len(s) {
			if _, ok := dict[s[end+1]]; ok {
				break
			} else {
				end++
				dict[s[end]] = end
			}
		}
		temp := end - i + 1
		if length < temp {
			length = temp
		}
	}
	return length
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	i := 0
	lookup := make(map[byte]int)
	max := 0
	for j := 0; j < n; j++ {
		lookup[s[j]]++
		for i <= j && lookup[s[j]] > 1 {
			lookup[s[i]]--
			i++
		}
		temp := j - i + 1
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bb"))       // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring(""))         // 0
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
}
