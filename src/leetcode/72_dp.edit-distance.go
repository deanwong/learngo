package main

import (
	"fmt"
)

// O(m*n)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// definition: dp[i][j] 表示要 word1[0:i-1] 变成 word1[0:j-1] 需要操作的最少字符数（步数）
	// formulation : word1[i-1] == word2[j-1] (即两个字符一样所以不需要操作) => dp[i][j] = dp[i-1][j-1]
	// word1[i-1] != word2[j-1]
	// 即两个字符不一样，
	// a) 要么取 word1[0:i-2] 与 word1[0:j-1](dp[i-1][j]) 需要操作的最少字符数，即删减掉word1的一个字符，+1
	// b) 要么取 word1[0:i-1] 与 word1[0:j-2](dp[i][j-1]) 需要操作的最少字符数，即删增加word1的一个字符，+1
	// c) 要么取 word1[0:i-2] 与 word1[0:j-2](dp[i-1][j-1]) 需要操作的最少字符数，即删改变word1的一个字符，+1
	// 如计算dp[2][2] 即 ho 和 ro 相等要操作几次，此时比较的字符正好相等，dp[1][1]是 h 变成 r 要操作 1 次，取dp[1][1]即可
	// 如计算dp[2][3] 是 ho 变成 ros 需要操作几次次，dp[2][2] 是 ho 和 ro 需要操作 1 次，dp[1][3]是 h 变成 ros 需要操作 3 次， dp[1][2]是 h 变成 ro 需要操作 2 次， 此时需要取三者最小值即dp[2][2]的操作再加 1
	// initial: dp[i][j] 当 i=0 或j=0 时都表示空字符，任何字符到要变成空字符串，都必须删除该字符的长度次
	// answer: dp[m][n]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if i == 0 && j > 0 {
				dp[i][j] = j
			} else if i > 0 && j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

func main() {
	/*
		+------+----+----+---+---+---+
		| text | i  | "" | r | o | s |
		+------+----+----+---+---+---+
		| j    | dp | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| ""   | 0  | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| h    | 1  | 1  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| o    | 2  | 2  | 2 | 1 | 2 |
		+------+----+----+---+---+---+
		| r    | 3  | 3  | 2 | 2 | 2 |
		+------+----+----+---+---+---+
		| s    | 4  | 4  | 3 | 3 | 2 |
		+------+----+----+---+---+---+
		| e    | 5  | 5  | 4 | 4 | 3 |
		+------+----+----+---+---+---+
	*/
	fmt.Println(minDistance("horse", "ros")) // 3
}
