package main

import (
	"fmt"
	"math"
	"sort"
)

// O(n*n)
func coinChange(coins []int, amount int) int {
	n := len(coins)
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// definition: dp[i] 表示要凑齐金额为 i 所要用的最少硬币数量
	// formulation: dp[i] = dp[i-coin[j]] + 1 j为物品栏所有选择 的最小值
	// 例如计算 dp[11]，需要取 dp[10]+1, dp[9]+1 和 dp[6]+1的最小值，地推的计算dp[6],需要取 dp[5]+1, dp[4]+1 和 dp[1]+1的最小值
	// initial: dp[0]=0
	// answer: dp[amount]
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	/*
		for i := 1; i <= amount; i++ {
			for j := 0; j < n; j++ {
				// 硬币面值小于等于背包总额并且扣掉硬币面值的前导dp已经计算过
				if coins[j] <= i && dp[i-coins[j]] != math.MaxInt64 {
					dp[i] = min(dp[i], dp[i-coins[j]]+1)
				}
			}
		}
	*/
	// 物品栏
	for j := 0; j < n; j++ {
		// 背包
		for i := coins[j]; i <= amount; i++ {
			if dp[i-coins[j]] != math.MaxInt64 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func coinChange_backtrace(coins []int, amount int) int {
	ans := math.MaxInt64
	n := len(coins)
	sort.Ints(coins)
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var backtrace func(targetLeft int, start int, combination []int)
	backtrace = func(targetLeft int, start int, combination []int) {
		if targetLeft == 0 {
			ans = min(ans, len(combination))
			return
		}
		for i := start; i < n && coins[i] <= targetLeft; i++ {
			combination = append(combination, coins[i])
			backtrace(targetLeft-coins[i], i, combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtrace(amount, 0, []int{})
	if ans == math.MaxInt64 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) // 3
	// 	fmt.Println(coinChange([]int{2}, 3))                                                             // -1
	// 	fmt.Println(coinChange([]int{1}, 0))                                                             // 0
	// 	fmt.Println(coinChange([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864)) // 24
}
