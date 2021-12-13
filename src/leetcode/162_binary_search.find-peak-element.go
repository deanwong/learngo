package main

import (
	"fmt"
	"math"
)

// O(logn)
func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i < 0 || i >= len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	lo, hi, mid := 0, len(nums)-1, 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		// peak
		if get(mid) > get(mid+1) && get(mid) > get(mid-1) {
			return mid
		}
		// decreasing, discard right part
		if nums[mid] > nums[mid+1] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return nums[mid]
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))          // 2
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) // 5
	fmt.Println(findPeakElement([]int{1}))                   // 0
}
