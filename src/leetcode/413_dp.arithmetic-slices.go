package main

import (
	"fmt"
)

// O(n)
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	ans := 0
	if n < 3 {
		return ans
	}
	// definition: dp[i] 表示以i为结尾的等差数列的数量
	// formulation: nums[i]-nums[i-1] == nums[i-1]-nums[i-2] => dp[i] = dp[i-1] + 1
	// initial: dp[0],dp[1] = 0;
	// answer: dp[0] + .. + dp[i] + .. + dp[n-1]
	dp := make([]int, n)
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			ans += dp[i]
		}
	}
	return ans
}
func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4})) // 3
	fmt.Println(numberOfArithmeticSlices([]int{1}))          // 0
}
