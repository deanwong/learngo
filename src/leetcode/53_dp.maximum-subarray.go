package main

import (
	"fmt"
)

// O(n)
func maxSubArray(nums []int) int {
	n := len(nums)
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// definition: dp[i] 前 i 个数的最大和
	// formulation : dp[i] = dp[i-1] + nums[i] 和 nums[i] 的最大值
	// initial: dp[0] = nums[0]
	dp := make([]int, n)
	ans := nums[0]
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5})) // 6
	fmt.Println(maxSubArray([]int{1}))                          // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))             // 23
}
