package main

import (
	"fmt"
)

// O(m+(n−m)×Σ)，其中 n 为字符串 s 的长度，m 为字符串 p 的长度，Σ 为所有可能的字符数
func findAnagrams(s string, p string) []int {
	letter, target := [26]int{}, [26]int{}
	ans := make([]int, 0)
	m, n := len(s), len(p)
	if m < n {
		return []int{}
	}
	for _, c := range p {
		target[c-'a']++
	}
	// left pointer
	i := 0
	for j := 0; j < m; j++ {
		letter[s[j]-'a']++
		fmt.Printf("%v:%v\n", string(s[j]), letter[s[j]-'a'])
		// move left pointer if letter count greater than target
		for i <= j && letter[s[i]-'a'] > target[s[i]-'a'] {
			letter[s[i]-'a']--
			fmt.Printf("discard %v left %v\n", string(s[i]), letter[s[i]-'a'])
			i++
		}
		if letter == target {
			// record left pointer
			fmt.Printf("found index %v\n", i)
			ans = append(ans, i)
		}
		fmt.Printf("i %v -> j %v\n", i, j)
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("abab", "ab"))        //[0,1,2]
}
