package main

import (
	"fmt"
)

// O(n)
func maxProfit(prices []int) int {
	n := len(prices)
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// definition: dp[i][0] 表示第 i 天没有股票； dp[i][1] 表示第 i 天持有股票
	// formulation : 分开讨论
	// dp[i][0] 可以是前一天也没有股票 dp[i-1][0] 或是前一天持有股票但是第 i 天卖出 dp[i-1][1] + prices[i]，取两者较大的值
	// dp[i][1] 可以是前一天也持有股票 dp[i-1][1] 或是前一天没有股票但是第 i 天买入 dp[i-1][0] - prices[i]，取两者较大的值
	// initial: dp[0] = 第一天的股价
	// answer: dp[n-1][0] 最后一天没有股票即maxProfit
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	fmt.Println(dp)
	return dp[n-1][0]
}

func maxProfit_greedy(prices []int) int {
	n := len(prices)
	// 连续上涨交易日，p1买pn卖收益最大，等价于 pn-p1 = p2-p1 + p3-p2 +...+ pn-pn-1
	// 连续下跌交易日不卖最赚
	ans := 0
	for i := 1; i < n; i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			ans += tmp
		}
	}
	return ans
}

func main() {
	/*
		+-------+-----------+-------------+---------------+
		| *     | has_stock | 0           | 1             |
		+-------+-----------+-------------+---------------+
		| price | *         | *           | *             |
		+-------+-----------+-------------+---------------+
		| 7     | [0]       | 0           | -7            |
		+-------+-----------+-------------+---------------+
		| 1     | [1]       | max(0,-6)=0 | max(-7,-1)=-1 |
		+-------+-----------+-------------+---------------+
		| 5     | [2]       | max(0,4)=4  | max(-1,-5)=-1 |
		+-------+-----------+-------------+---------------+
		| 3     | [3]       | max(4,3)=4  | max(-1,1)=1   |
		+-------+-----------+-------------+---------------+
		| 6     | [4]       | max(4,7)=7  | max(1,-2)=1   |
		+-------+-----------+-------------+---------------+
		| 4     | [5]       | max(7,5)=7  | max(1,3)=3    |
		+-------+-----------+-------------+---------------+
	*/
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))        // 7
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))           // 4
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))           // 0
	fmt.Println(maxProfit_greedy([]int{7, 1, 5, 3, 6, 4})) // 7
	fmt.Println(maxProfit_greedy([]int{1, 2, 3, 4, 5}))    // 4
	fmt.Println(maxProfit_greedy([]int{7, 6, 4, 3, 1}))    // 0
}
