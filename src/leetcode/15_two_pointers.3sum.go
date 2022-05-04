package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	twoSum := func(l int, r int, target int, value int) [][]int {
		var res = make([][]int, 0)
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[l], nums[r], value})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
		return res
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		ans = append(ans, twoSum(i+1, n-1, -nums[i], nums[i])...)
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{}))                    // []
	fmt.Println(threeSum([]int{0}))                   // []
}
