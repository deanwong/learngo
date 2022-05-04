package main

import (
	"fmt"
)

// O(n^2)
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// definition: dp[i] 表示 i 之前包括 i 的最长上升子序列的长度
	// definition(2): count[i] 表示 i 之前包括 i 的最长上升子序列的数量
	// formulation : nums[i] > nums[j] => dp[i] = max(dp[j]+1) 0<=j<i；
	// formulation(2) : dp[j]+1 > dp[i] => 表示发现更长的子序列，更新count[i]为count[j], dp[j]+1 == dp[i] => 表示发现了两个一样长的子序列，累加
	// initial: dp[i] = 1, 至少有一个上升子串; count[i] = 1, 子序列数据量至少也是1
	// answer: max(dp[i]) 先得到最长上升子序列的长度，再遍历dp比对长度，一旦符合累计 i 对应的 count
	dp := make([]int, n)
	count := make([]int, n)
	dp[0] = 1
	count[0] = 1
	maxLen := dp[0]
	ans := 0
	for i := 1; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					count[i] = count[j]
					dp[i] = dp[j] + 1
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	for i := 0; i < n; i++ {
		if dp[i] == maxLen {
			ans += count[i]
		}
	}
	fmt.Println(maxLen)
	fmt.Println(dp)
	fmt.Println(count)
	return ans
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7})) // 2
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2})) // 5
	fmt.Println(findNumberOfLIS([]int{1}))             // 1
}
