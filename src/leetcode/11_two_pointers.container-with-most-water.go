package main

import "fmt"

// O(n)
func maxArea(height []int) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		distance := right - left
		ans = max(ans, distance*min(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             // 16
	fmt.Println(maxArea([]int{1, 2, 1}))                   // 12
}
