package main

import (
	"fmt"
)

// O(n^2) n为字符串长度
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	// definition: dp[i][j] 表示i为起点j为终点的子串是否是回文
	// formulation: nums[i] == nums[j] => dp[i][j] = dp[i+1][j-1]
	// j-1 - (i+1) +1 < 2 ==> j - i < 3（i和j中相距0~1个字符时) AND nums[i] == nums[j] => dp[i][j] = true
	// initial: dp[i][i] = true
	// answer: dp[i][j]时取最长子串长度
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	maxLen, begin := 1, 0
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}
			if j-i < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			len := j - i + 1
			if dp[i][j] && len > maxLen {
				maxLen = len
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("babad")) // bab
	fmt.Println(longestPalindrome("cbbd"))  // bb
	fmt.Println(longestPalindrome("ac"))    // a
	fmt.Println(longestPalindrome("aa"))    // aa
}
