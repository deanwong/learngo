package main

import (
	"fmt"
)

// O(n)
func numDecodings(s string) int {
	n := len(s)
	// definition: dp[i] 表示以i为结尾的字符串的组合数
	// formulation:
	// 1. s[i]>=1 && s[i]<=9 只能有位置 i 组成一种方式 dp[i] = dp[i-1]
	// 2. s[i-1] *10 + s[i]>=10 && s[i-1] *10 + s[i]<=26 只能当前位置 i 的与前一位置（i-1）组成 dp[i] = dp[i-2]
	// 要么是1要么是2，所以状态转移方程是 dp[i] = dp[i-1] + dp[i-2]
	// initial: dp[0] = 1 单个字符只要不为‘0’就有一种组合
	// answer: dp[n]
	dp := make([]int, n+1)
	s = " " + s
	dp[0] = 1
	for i := 1; i <= n; i++ {
		a := s[i] - '0'
		b := (s[i-1]-'0')*10 + (s[i] - '0')
		// 如果 a 属于有效值，那么 dp[i] 可以由 dp[i - 1] 转移过来
		if a >= 1 && a <= 9 {
			dp[i] = dp[i-1]
		}
		// 如果 b 属于有效值，那么 dp[i] 可以由 dp[i - 2] 或者 dp[i - 1] + dp[i - 2] 转移过来
		// dp[i] 此时已经是从 dp[i-1] 转移而来
		if b >= 10 && b <= 26 {
			dp[i] = dp[i-2] + dp[i]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numDecodings("12"))  // 2
	fmt.Println(numDecodings("226")) // 3
	fmt.Println(numDecodings("06"))  // 0
}
