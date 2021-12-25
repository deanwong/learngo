package main

import "fmt"

// O(n)
func canJump(nums []int) bool {
	n := len(nums)
	fartest := nums[0]
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		// 当前位置不可达直接返回
		if fartest < i {
			return false
		}
		// 最远可达已经超过最后一位直接返回
		if fartest >= n-1 {
			return true
		}
		// 可达，更新最远可达距离
		fartest = max(nums[i]+i, fartest)
		//fmt.Println(fartest)
	}
	return true
}

// dp O(n*n)
func canJump2(nums []int) bool {
	n := len(nums)
	// definition: dp[i] 表示i节点是否可达
	// formulation: 如果 j+nums[j] >= i 即 j 可达 i && dp[j] = true => dp[i] = true
	// initial: dp[0] = true
	// answer: dp[n-1]
	dp := make([]bool, n)
	if n == 0 {
		return false
	}
	// 首节点肯定可达
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 前面的节点可达且前面节点可达节点超过i的位置
			if dp[j] && j+nums[j] >= i {
				dp[i] = true
			}
		}
	}
	return dp[n-1]
}

// 倒序
func canJump3(nums []int) bool {
	n := len(nums)
	last := n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= last {
			last = i
		}
	}
	return last == 0
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	fmt.Println(canJump([]int{0}))             // false
}
