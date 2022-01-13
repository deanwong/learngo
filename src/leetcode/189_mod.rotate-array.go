package main

import "fmt"

// O(n)
func rotate(nums []int, k int) {
	n := len(nums)
	// 整体向右偏移的offset
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		// 滚动数组
		ans[(i+k)%n] = nums[i]
	}
	copy(nums, ans)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums) // [5,6,7,1,2,3,4]

}
