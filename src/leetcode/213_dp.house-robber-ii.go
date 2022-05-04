package main

import "fmt"

// O(n)
func rob(nums []int) int {
	n := len(nums)
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	_rob := func(houses []int) int {
		m := len(houses)
		if m <= 0 {
			return 0
		}
		if m == 1 {
			return houses[0]
		}
		dp := make([]int, m)
		dp[0] = houses[0]
		dp[1] = max(houses[0], houses[1])
		for i := 2; i < m; i++ {
			dp[i] = max(dp[i-2]+houses[i], dp[i-1])
		}
		return dp[m-1]
	}
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// 分开讨论，一种是偷第一家不偷最后一家，另一种是不偷第一家偷最后一家
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))    // 3
	fmt.Println(rob([]int{1, 2, 3, 1})) // 4
	fmt.Println(rob([]int{1, 2, 3}))    // 3
	fmt.Println(rob([]int{1}))          // 1
}
