package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	twosum := func(left int, right int, target int, value int) [][]int {
		res := make([][]int, 0)
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{value, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
			// fmt.Printf("left %v, right %v, target %v\n", nums[left], nums[right], target)
		}
		return res
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ans = append(ans, twosum(i+1, n-1, -nums[i], nums[i])...)
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{}))                    // []
	fmt.Println(threeSum([]int{0}))                   // []
}
