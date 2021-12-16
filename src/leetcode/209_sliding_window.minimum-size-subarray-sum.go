package main

import (
	"fmt"
	"math"
)

// O(n)
func minSubArrayLen(target int, nums []int) int {
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// left pointer
	i := 0
	ans := math.MaxInt64
	n := len(nums)
	if n <= 0 {
		return 0
	}
	sum := 0
	// right pointer
	for j := 0; j < n; j++ {
		sum += nums[j]
		// if >= target try to move left pointer to find min length subarray
		for i <= j && sum >= target {
			// fmt.Printf("i %v -> j %v len %v\n", i, j, j-i+1)
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if ans == math.MaxInt64 {
		ans = 0
	}
	return ans
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))          // 5
}
