package main

import (
	"fmt"
	"sort"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	res := ""
	for i := 0; i < len(nums); i++ {
		res += fmt.Sprintf("%d", nums[i])
	}
	return res
}

func main() {
	fmt.Println(minNumber([]int{10, 2}))           // 102
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9})) // 3033459
}
