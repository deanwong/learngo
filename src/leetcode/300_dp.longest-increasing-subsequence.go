package main

import (
	"fmt"
)

// O(n^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	// definition: dp[i] 表示 i 之前包括 i 的最长上升子序列的长度
	// formulation : nums[i] > nums[j] => dp[i] = max(dp[j]+1) 0<=j<i
	// initial: dp[i] = 1, 至少有一个上升子串
	// answer: max(dp[i])
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n)
	dp[0] = 1
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// fmt.Println(dp)
		ans = max(ans, dp[i])
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           // 4
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))        // 1
	fmt.Println(lengthOfLIS([]int{0}))                          // 1
}
