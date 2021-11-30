package main

import "fmt"

// O(logn)
func search(nums []int, target int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		// 有序的就可以比较，无序的就继续二分直到变成有序的
		if nums[lo] <= nums[mid] {
			// 左边有序
			if nums[lo] <= target && target < nums[mid] {
				// target在有序的一边
				hi = mid - 1
			} else {
				// target在无序的一边
				lo = mid + 1
			}
		} else {
			// 右边有序
			if nums[mid] < target && target <= nums[hi] {
				// target在有序的一边
				lo = mid + 1
			} else {
				// target在无序的一边
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}, 7))

}
