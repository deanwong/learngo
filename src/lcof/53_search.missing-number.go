package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] == mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))                   // 2
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9})) // 8
}
