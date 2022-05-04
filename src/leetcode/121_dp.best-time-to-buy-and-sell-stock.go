package main

import (
	"fmt"
)

// O(n)
func maxProfit(prices []int) int {
	n := len(prices)
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	// definition: dp[i] 表示前 i 天股票的最低价
	// formulation : dp[i] = min(dp[i-1], 当天的股价)
	// initial: dp[0] = 第一天的股价
	// answer: 使用 maxProfit 变量记录 max(当天的股价-dp[i])
	dp := make([]int, n)
	dp[0] = prices[0]
	for i := 1; i < n; i++ {
		dp[i] = min(prices[i], dp[i-1])
		ans = max(ans, prices[i]-dp[i])
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
