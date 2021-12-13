package main

import "fmt"

// O(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([][2]int, len(nums))
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func rob_p(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		// max(f(n-2) + tonight, f(n-1) + 0)
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func rob_plus(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	a := nums[0]
	b := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		temp := b
		b = max(a+nums[i], b)
		a = temp
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob_p([]int{1, 2, 3, 1}))
	fmt.Println(rob_p([]int{2, 7, 9, 3, 1}))

}
