package main

import "fmt"

func maxSubArray(nums []int) int {
	// definition: dp[i] 表示前 i 个数字（实际是i-1）组成数据中最大的子数组的和
	// formulation : dp[i] = max(nums[i], dp[i-1]+nums[i])
	// initial: dp[0] 第 1 个数字
	// answer: 使用变量记录 max
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if n == 0 {
		return 0
	}
	ans := nums[0]
	a := nums[0]
	for i := 1; i < n; i++ {
		b := max(nums[i], a+nums[i])
		ans = max(ans, b)
		a = b
		b = ans
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
}
