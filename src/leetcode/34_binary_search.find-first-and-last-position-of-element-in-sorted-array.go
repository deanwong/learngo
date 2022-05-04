package main

import (
	"fmt"
)

// O(logn)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	fisrt := findFirst(nums, target)
	if fisrt == -1 {
		return []int{-1, -1}
	}
	last := findLast(nums, target)
	return []int{fisrt, last}
}

func findFirst(nums []int, target int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo < hi {
		mid = lo + (hi-lo)/2
		if nums[mid] < target {
			lo = lo + 1
		} else if nums[mid] == target {
			hi = mid
		} else {
			hi = mid - 1
		}
	}
	if nums[lo] != target {
		return -1
	}
	return lo
}

func findLast(nums []int, target int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo < hi {
		mid = lo + (hi-lo)/2 + 1
		if nums[mid] < target {
			lo = lo + 1
		} else if nums[mid] == target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{2, 2}, 2))
}
