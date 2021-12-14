package main

import "fmt"

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
	ans := 0
	dict := make(map[byte]int)
	i := 0
	for j := 0; j < len(s); j++ {
		// move right pointer and add
		dict[s[j]]++
		for i <= j && dict[s[j]] > 1 {
			// move left if find duplicate char
			dict[s[i]]--
			i++
		}
		temp := j - i + 1
		if ans < temp {
			ans = temp
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
