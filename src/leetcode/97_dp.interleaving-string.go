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

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if m+n != len(s3) {
		return false
	}
	// definition: dp[i][j] 表示 s1[0~i-1]和s2[0~j-1]能否交错组成s3[0~i+j-1]
	// formulation: dp[i][j] = (dp[i-1][j] AND s1[i-1] = s3[i+j-1]) OR (dp[i][j-1] AND s2[j-1] = s3[i+j-1])
	// s1 的前 i-1 个字符和 s2 的前 j 个字符能否构成 s3 的前 i+j−1 位 且 s1的i个字符等于 s3的 i+j 个字符
	// initial:
	// case 1：dp[0][0] = true
	// case 2：if j=0 : dp[i][0] = dp[i-1][0] AND s1[i-1] = s3[i-1] 表示 s1 的前 i 位是否能构成 s3 的前 i 位。因此需要满足的条件为，前 i-1 位可以构成 s3 的前 i-1 位且 s1 的第 i 位（s1[i-1]）等于 s3 的第 i 位（s3[i-1]）
	// case 3：if i=0 : dp[0][j] = dp[0][j-1] AND s2[j-1] = s3[j-1] 表示 s2 的前 j 位是否能构成 s3 的前 j 位。因此需要满足的条件为，前 j-1 位可以构成 s3 的前 j-1 位且 s3 的第 j 位（s2[j-1]）等于 s3 的第 j 位（s3[j-1]）
	// answer: dp[m][n]
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

func main() {
	/*
		如果上一行相同列是 T ，那么看左边字符
		如果左一列相同行是 T ，那么看上边字符
		+------+----+-----+---+---+---+---+---+
		| text | i  | ""  | d | b | b | c | a |
		+------+----+-----+---+---+---+---+---+
		| j    | dp | 0   | 1 | 2 | 3 | 4 | 5 |
		+------+----+-----+---+---+---+---+---+
		| ""   | 0  | T   | F | F | F | F | F |
		+------+----+-----+---+---+---+---+---+
		| a    | 1  | T   | F | F | F | F | F |
		+------+----+-----+---+---+---+---+---+
		| a    | 2  | T   | T | T | T | T | F |
		+------+----+-----+---+---+---+---+---+
		| b    | 3  | F   | T | T | F | T | F |
		+------+----+-----+---+---+---+---+---+
		| c    | 4  | F   | F | T | T | T | T |
		+------+----+-----+---+---+---+---+---+
		| c    | 5  | F   | F | F | T | F | T |
		+------+----+-----+---+---+---+---+---+
	*/
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false
	fmt.Println(isInterleave("", "", ""))                     // true
}
