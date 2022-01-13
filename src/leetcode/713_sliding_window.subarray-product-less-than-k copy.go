package main

import (
	"fmt"
)

// O(n)
func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	temp := 1
	if k <= 1 {
		return 0
	}
	// left pointer
	i := 0
	// right pointer
	for j := 0; j < len(nums); j++ {
		temp *= nums[j]
		for i <= j && temp >= k {
			temp /= nums[i]
			fmt.Printf("discard %v left %v\n", nums[i], temp)
			i++
		}
		// key point
		ans += j - i + 1
		fmt.Printf("i %v -> j %v add %v\n", i, j, j-i+1)
	}
	return ans
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)) // 8
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))       // 0
	fmt.Println(numSubarrayProductLessThanK([]int{1, 1, 1}, 2))       // 6
}
