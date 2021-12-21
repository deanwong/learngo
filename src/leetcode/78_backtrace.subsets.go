package main

import "fmt"

// O(n * 2^n) 一共 2^n 个状态，每种状态需要 O(n)的时间来构造子集
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	var backtrace func(start int, subset []int)
	backtrace = func(start int, subset []int) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		ans = append(ans, subsetCopy)
		for i := start; i < n; i++ {
			subset = append(subset, nums[i])
			// fmt.Printf("before idx %v, subset %v\n", start, subset)
			backtrace(i+1, subset)
			subset = subset[:len(subset)-1]
			// fmt.Printf("after idx %v, subset %v\n", start, subset)
		}
	}
	backtrace(0, []int{})
	return ans
}

func subsets2(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	var backtrace func(start int, setSize int, subset []int)
	backtrace = func(start int, setSize int, subset []int) {
		if setSize == 0 {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			ans = append(ans, subsetCopy)
			return
		}
		for i := start; i < n; i++ {
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
	fmt.Println(subsets([]int{1, 2, 3}))  // [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	fmt.Println(subsets2([]int{1, 2, 3})) // [[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]]
}
