package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	zero := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zero++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	// 最大牌 - 最小牌 < 5 则可构成顺子
	return nums[4]-nums[zero] < 5
}

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5})) // true
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5})) // true
}
