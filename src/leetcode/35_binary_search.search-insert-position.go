package main

import "fmt"

// O(logn)
func searchInsert(nums []int, target int) int {
	lo, hi, mid := 0, len(nums), 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if mid >= len(nums) || nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

}
