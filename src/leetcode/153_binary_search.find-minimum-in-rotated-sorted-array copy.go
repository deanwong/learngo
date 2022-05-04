package main

import "fmt"

// O(logn)
func findMin(nums []int) int {
	return binarySearchan(nums, 0, len(nums)-1, nums[0])
}

func findMin2(nums []int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo < hi {
		mid = lo + (hi-lo)/2
		if nums[mid] <= nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return nums[lo]
}

func binarySearchan(nums []int, lo int, hi int, ans int) int {
	if lo > hi {
		return ans
	}
	mid := lo + (hi-lo)/2
	if nums[lo] <= nums[mid] {
		ans = min(ans, nums[lo])
		lo = mid + 1
		// fmt.Printf("lo %v hi %v, ans %v\n", lo, hi, ans)
		return binarySearchan(nums, lo, hi, ans)
	} else {
		ans = min(ans, nums[mid])
		hi = mid - 1
		// fmt.Printf("lo %v hi %v, ans %v\n", lo, hi, ans)
		return binarySearchan(nums, lo, hi, ans)
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{2, 1}))

	fmt.Println(findMin2([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin2([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin2([]int{11, 13, 15, 17}))
	fmt.Println(findMin2([]int{2, 1}))

}
