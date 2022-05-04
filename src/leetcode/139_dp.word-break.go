package main

import (
	"fmt"
)

// O(n^2)
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	// definition: dp[i] 表示长度为 i 的字符串是否在字典中存在
	// 使用 j 分割 0:i 的子串，j从i-1开始往前遍历
	// dp[i]=>s[0:i-1] 是否为真 取决于 dp[j]=>s[0:j-1] 是否为真 AND 剩余子串 s[j:i-1] (go中表示为s[j:i]) 是否存在于词典中
	// formulation: dp[i] = s[j:i]存在于字典 && dp[j] 0<=j<i
	// initial: dp[0] = true
	// answer: dp[n]
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			// fmt.Printf("i %v, j %v, substr %v, wordMap[s[j:i]] %v\n", i, j, s[j:i], wordMap[s[j:i]])
			if wordMap[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))                       // true
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))                  // true
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}
