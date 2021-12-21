package main

import (
	"fmt"
	"sort"
)

// O(n * 2^n) 一共 2^n 个状态，每种状态需要 O(n)的时间来构造子集
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	var backtrace func(start int, subset []int)
	backtrace = func(start int, subset []int) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		ans = append(ans, subsetCopy)
		for i := start; i < n; i++ {
			// 去重
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			subset = append(subset, nums[i])
			backtrace(i+1, subset)
			subset = subset[:len(subset)-1]
		}
	}
	backtrace(0, []int{})
	return ans
}

func subsetsWithDup2(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	var backtrace func(start int, setSize int, subset []int)
	backtrace = func(start int, setSize int, subset []int) {
		if setSize == 0 {
			//subset元素数量
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			ans = append(ans, subsetCopy)
			return
		}
		for i := start; i < n; i++ {
			// 去重
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			subset = append(subset, nums[i])
			backtrace(i+1, setSize-1, subset)
			subset = subset[:len(subset)-1]
		}
	}
	for k := 0; k <= len(nums); k++ {
		backtrace(0, k, []int{})
	}
	return ans
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))  // [[],[1],[1,2],[1,2,2],[2],[2,2]]
	fmt.Println(subsetsWithDup2([]int{1, 2, 2})) // [[] [1] [2] [1 2] [2 2] [1 2 2]]
}
