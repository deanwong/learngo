package main

import "fmt"

func maxProfit(prices []int) int {
	// definition: dp[i] 表示前 i 天（实际是第i-1天）股票的最低价
	// formulation : dp[i] = min(dp[i-1], 当天的股价)
	// initial: dp[0] = 第一天的股价
	// answer: 使用 maxProfit 变量记录 max(当天的股价-dp[i])
	n := len(prices)
	dp := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	ans := 0
	if n == 0 {
		return ans
	}
	dp[0] = prices[0]
	for i := 1; i < n; i++ {
		dp[i] = min(dp[i-1], prices[i])
		ans = max(prices[i]-dp[i], ans)
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
