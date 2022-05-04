package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)
	i, j := 0, n-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{nums[i], nums[j]}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [2,7]
}
