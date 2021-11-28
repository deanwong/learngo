package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
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

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
