package main

import (
	"fmt"
	"math"
)

// O(n*n)
func numSquares(n int) int {
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// definition: dp[i] 表示和为i的完全平方数的最少数量
	// formulation: dp[i] = dp[i-j*j] + 1 j为物品栏所有选择 中的最小值
	// 例如计算 dp[4]，需要取 dp[3]+1, dp[2]+1 和 dp[1]+1的最小值
	// initial: dp[0]=0
	// answer: dp[n]
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	// 物品栏
	for j := 1; j*j <= n; j++ {
		// 背包
		for i := j * j; i <= n; i++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}
	/*
		// 背包
		for i := 1; i <= n; i++ {
			// 物品栏
			for j := 1; j*j <= i; j++ {
				dp[i] = min(dp[i-j*j]+1, dp[i])
			}
		}
	*/
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12)) // 3
	fmt.Println(numSquares(13)) // 2
}
