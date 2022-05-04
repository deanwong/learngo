package main

import "fmt"

func exchange(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	i, j := 0, n-1
	for i < j {
		if nums[i]%2 == 0 && nums[j]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else if nums[i]%2 == 0 {
			j--
		} else {
			i++
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4})) // [1,3,2,4]
}
