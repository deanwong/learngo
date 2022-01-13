package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	lookup := make(map[int]int)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			lookup[num] = stack[len(stack)-1]
		} else {
			lookup[num] = -1
		}
		stack = append(stack, num)
	}
	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = lookup[nums1[i]]
	}
	return ans
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})) // [-1,3,-1]
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))    // [3,-1]
}
