package main

import (
	"fmt"
	"sort"
)

// O(d×(m+n))，其中 d 表示 dictionary 的长度，m 表示 s 的长度，n 表示 dictionary 中字符串的平均长度。我们需要遍历 dictionary 中的 d 个字符串，每个字符串需要 O(n+m) 的时间复杂度来判断该字符串是否为 s 的子序列
func findLongestWord(s string, dictionary []string) string {
	n := len(s)
	sort.Strings(dictionary)
	ans := ""
	find := func(s2 string) string {
		m := len(s2)
		j := 0
		for i := 0; i < n; i++ {
			if j < m && s[i] == s2[j] {
				j++
			}
		}
		if j == m {
			return s2
		}
		return ""
	}
	for _, s2 := range dictionary {
		res := find(s2)
		if len(res) > len(ans) {
			ans = res
		}
	}
	return ans
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))                                                 // apple
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))                                                                    // a
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea", "abpcplaaa", "abpcllllll", "abccclllpppeeaaaa"})) // apple
	fmt.Println(findLongestWord("abce", []string{"abe", "abc"}))                                                                         // abc
}
