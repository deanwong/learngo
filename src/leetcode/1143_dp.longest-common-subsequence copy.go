package main

import (
	"fmt"
)

// O(m*n)
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// definition: dp[i][j] 表示 text1[0:i-1] 和 text2[0:j-1] 的公共子序列的长度
	// formulation : text1[i-1] == text2[j-1] (即两个字符一样所以必然在公共子序列中) => dp[i][j] = 1 + dp[i-1][j-1]
	// text1[i-1] != text2[j-1] (即两个字符至少有一个不在公共子序列中,两个都不在的场景不用考虑，因为两个都不在时长度肯定比其中一个在的长度要短) => dp[i][j] = max(dp[i-1][j],dp[i][j-1])
	// initial: dp[i][j] 当 i=0 或j=0 时都表示空字符，此时肯定没有字符在公共子序列中，所以为 0
	// answer: dp[m][n]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	//fmt.Println(dp)
	return dp[m][n]
}

func main() {
	/*
		+------+----+----+---+---+---+
		| text | i  | "" | a | c | e |
		+------+----+----+---+---+---+
		| j    | dp | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| ""   | 0  | 0  | 0 | 0 | 0 |
		+------+----+----+---+---+---+
		| a    | 1  | 0  | 1 | 1 | 1 |
		+------+----+----+---+---+---+
		| b    | 2  | 0  | 1 | 1 | 1 |
		+------+----+----+---+---+---+
		| c    | 3  | 0  | 1 | 2 | 2 |
		+------+----+----+---+---+---+
		| d    | 4  | 0  | 1 | 2 | 2 |
		+------+----+----+---+---+---+
		| e    | 5  | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
	*/
	fmt.Println(longestCommonSubsequence("abcde", "ace")) // 3
	/*
		+------+----+----+---+---+---+
		| text | i  | "" | a | b | c |
		+------+----+----+---+---+---+
		| j    | dp | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| ""   | 0  | 0  | 0 | 0 | 0 |
		+------+----+----+---+---+---+
		| a    | 1  | 0  | 1 | 1 | 1 |
		+------+----+----+---+---+---+
		| b    | 2  | 0  | 1 | 2 | 2 |
		+------+----+----+---+---+---+
		| c    | 3  | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
	*/
	fmt.Println(longestCommonSubsequence("abc", "abc")) // 3
	fmt.Println(longestCommonSubsequence("abc", "def")) // 0
}
