package main

import "fmt"

func findRepeatNumber(nums []int) int {
	dict := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := dict[nums[i]]; ok {
			return nums[i]
		}
		dict[nums[i]] = nums[i]
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		// 数字已在对应索引位置，无需交换
		if nums[i] == i {
			i++
			continue
		}
		// 代表索引 nums[i] 处和索引 i 处的元素一样，表示有重复值
		// 如当前索引 i=4 nums[i]=2, 并且 nums[nums[i]] 即 nums[2]=2，都等于2，表示2是重复值
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		// 交换
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3})) // 2 or 3
	fmt.Println(findRepeatNumber2([]int{3, 4, 2, 0, 0, 1}))    // 2 or 3
}
