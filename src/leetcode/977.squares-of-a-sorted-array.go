package main

import "fmt"

// O(n)
func sortedSquares(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	left, right := 0, length-1
	for i := length - 1; left <= right; i-- {
		l, r := nums[left]*nums[left], nums[right]*nums[right]
		if l > r {
			ans[i] = l
			left++
		} else {
			ans[i] = r
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
