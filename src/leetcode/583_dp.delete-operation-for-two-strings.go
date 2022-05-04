package main

import (
	"fmt"
)

// O(m*n)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// definition: dp[i][j] 表示要 word1[0:i-1] 与 word1[0:j-1] 需要删除的最少的字符数
	// formulation : word1[i-1] == word2[j-1] (即两个字符一样所以不需要删除任一字符) => dp[i][j] = dp[i-1][j-1]
	// word1[i-1] != word2[j-1]  (即两个字符不一样，要么取 word1[0:i-2] 与 word1[0:j-1](dp[i-1][j])需要删除的最少的字符数， 要么取 word1[0:i-1] 与 word1[0:j-2](dp[i][j-1])需要删除的最少的字符数，取更小者再加 1)
	// 如计算dp[2][2] 即 se 和 ea 相等要删除几个字符，dp[2][1] 是 se 和 e 需要删除 1 个字符，dp[1][2] 是 s 和 ea 需要删除 3 个字符，此时取较小前者 1，另外由于 e 和 a 不等，所以还要删除其中一个字符（这里选的是前者所以会删除 word2的 a）需要 + 1
	// initial: dp[i][j] 当 i=0 或j=0 时都表示空字符，任何字符到要变成空字符串，都必须删除该字符的长度次
	// answer: dp[m][n]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				// 空字符和空字符不用删除任何字符
				dp[i][j] = 0
			} else if i == 0 && j > 0 {
				// word[0:j-1]要变成空字符需要，删除j个字符
				dp[i][j] = j
			} else if i > 0 && j == 0 {
				// word[0:i-1]要变成空字符需要，删除i个字符
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
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
		| text | i  | "" | e | a | t |
		+------+----+----+---+---+---+
		| j    | dp | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| ""   | 0  | 0  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| s    | 1  | 1  | 2 | 3 | 4 |
		+------+----+----+---+---+---+
		| e    | 2  | 2  | 1 | 2 | 3 |
		+------+----+----+---+---+---+
		| a    | 3  | 3  | 2 | 1 | 2 |
		+------+----+----+---+---+---+
	*/
	fmt.Println(minDistance("sea", "eat")) // 2
}
