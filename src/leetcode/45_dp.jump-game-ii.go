package main

import (
	"fmt"
	"math"
)

// O(n)
// Greedy 难理解
func jump(nums []int) int {
	n := len(nums)
	steps, fartest, end := 0, 0, 0
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 不访问最后一个元素，这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置
	for i := 0; i < n-1; i++ {
		// 获取可达最远的下标
		fartest = max(nums[i]+i, fartest)
		// 如果达到上次的边界(最远的下标)，扩展边界，并记步
		// 边界的作用就是在迭代的过程中不断从 [i,end] 中获取能走的最远的下标
		if i == end {
			end = fartest
			steps++
		}
	}
	return steps
}

// dp
func jump2(nums []int) int {
	n := len(nums)
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// definition: dp[i] 表示到 i 节点最小步数
	// formulation: 如果 j+nums[j] >= i 即 j 可达 i => dp[i] = min(dp[j]+1) 0<=j<i
	// initial: dp[0] = 0
	// answer: dp[n-1]
	dp := make([]int, n)
	if n == 0 {
		return 0
	}
	// 首节点需要0步
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			// 逐个计算i之前的节点到i的最小步数
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(jump2([]int{2, 3, 1, 1, 4})) // 2
	fmt.Println(jump2([]int{2, 3, 0, 1, 4})) // 2
}
