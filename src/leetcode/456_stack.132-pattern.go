package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	n := len(nums)
	stack := []int{nums[n-1]}
	maxK := math.MinInt64
	fmt.Println(stack)
	for i := n - 2; i >= 0; i-- {
		num := nums[i]
		if num < maxK {
			return true
		}
		for len(stack) > 0 && num > stack[len(stack)-1] {
			// 此时栈内元素都可以作为2，并且栈底的是最大的
			maxK = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		fmt.Println(maxK)
		// if num > maxK {
		stack = append(stack, num)
		// }
		fmt.Println(stack)
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))    // false
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))    // true
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))   // true
	fmt.Println(find132pattern([]int{3, 5, 0, 3, 4})) // true
}
