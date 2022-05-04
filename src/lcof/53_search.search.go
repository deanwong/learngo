package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	n := len(nums)
	binarysearch := func(lo, hi int, val int) int {
		for lo <= hi {
			mid := lo + (hi-lo)/2
			// 严格小于才能保证加一的时候 mid 正好等于 target， 所以要取 lo
			if nums[mid] < val {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return lo
	}
	// 第一个等于 target 的位置
	l := binarysearch(0, n-1, target)
	// 第一个大于 target 的位置
	r := binarysearch(l, n-1, target+1)
	return r - l
}

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8)) // 2
	fmt.Println(search([]int{3, 4, 2, 0, 0, 1}, 6))  // 0
	fmt.Println(search([]int{2, 2}, 2))              // 2
	fmt.Println(search([]int{1}, 0))                 // 0
}
