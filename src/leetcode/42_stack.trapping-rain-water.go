package main

import (
	"fmt"
)

func trap(height []int) int {
	n := len(height)
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// 单调栈, 存的是下标
	stack := make([]int, 0)
	ans := 0
	for i := 0; i < n; i++ {
		// 当前元素大于栈顶元素，需要重新维护单调栈
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			fmt.Printf("i %v => value %v, top %v => value %v\n", i, height[i], top, height[top])
			fmt.Printf("before pop %v\n", stack)
			// pop
			stack = stack[:len(stack)-1]
			// 如果栈内还有和栈顶top一样的情况，继续pop
			for len(stack) > 0 && height[stack[len(stack)-1]] == height[top] {
				stack = stack[:len(stack)-1]
			}
			fmt.Printf("after pop %v\n", stack)

			if len(stack) > 0 {
				left := stack[len(stack)-1]
				// left 此时指向的是此次接住的雨水的左边界的位置。右边界是当前的柱体，即i。
				// min(height[left], height[i]) 是左右柱子高度的min，减去height[top]就是雨水的高度。
				// i - left - 1 是雨水的宽度
				fmt.Printf("height %v, width%v\n", min(height[left], height[i])-height[top], i-left-1)
				ans += (min(height[left], height[i]) - height[top]) * (i - left - 1)
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
}
