package main

import "fmt"

func countMaxOrSubsets(nums []int) int {
	ans := 0
	max := 0
	n := len(nums)
	var backtrace func(idx, orVal int)
	backtrace = func(idx, orVal int) {
		if idx == n {
			// 清空重算
			if orVal > max {
				max = orVal
				ans = 1
			} else if orVal == max {
				ans++
			}
			return
		}
		// 选择 nums[i]
		backtrace(idx+1, orVal|nums[idx])
		// 不选 nums[i]
		backtrace(idx+1, orVal)
	}
	backtrace(0, 0)
	return ans
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))       // 2
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))    // 7
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5})) // 6
}
