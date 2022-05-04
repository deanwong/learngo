package main

import (
	"fmt"
)

// O(logn)
func search(nums []int, target int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		// 左边有序
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// O(logn)
func search2(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, lo int, hi int, target int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		return mid
	}
	if nums[lo] == target {
		return lo
	}
	if nums[hi] == target {
		return hi
	}
	// 左边有序
	if nums[lo] < nums[mid] {
		if nums[lo] < target && target < nums[mid] {
			hi = mid - 1
			return binarySearch(nums, lo+1, hi, target)
		} else {
			lo = mid + 1
			return binarySearch(nums, lo, hi-1, target)
		}
	} else {
		if nums[mid] < target && target < nums[hi] {
			lo = mid + 1
			return binarySearch(nums, lo, hi-1, target)
		} else {
			hi = mid - 1
			return binarySearch(nums, lo+1, hi, target)
		}
	}
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}, 7))
	fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{1}, 0))

	fmt.Println(search2([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}, 7))
	fmt.Println(search2([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}, 3))
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search2([]int{1}, 0))
}
