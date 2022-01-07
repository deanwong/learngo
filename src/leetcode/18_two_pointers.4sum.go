package main

import (
	"fmt"
	"sort"
)

// O(n^3)
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	twoSum := func(left int, right int, target int, n2 int, n1 int) [][]int {
		res := make([][]int, 0)
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				fmt.Printf("left %v, right %v, sum %v\n", nums[left], nums[right], nums)
				res = append(res, []int{n1, n2, nums[left], nums[right]})
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
		}
		return res
	}
	ans := make([][]int, 0)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			ans = append(ans, twoSum(j+1, n-1, target-nums[i]-nums[j], nums[j], nums[i])...)
		}
	}
	return ans
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // [[2,2,2,2]]
	fmt.Println(fourSum([]int{0}, 0))                  // []
}
