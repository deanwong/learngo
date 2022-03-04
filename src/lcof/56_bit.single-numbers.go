package main

import (
	"fmt"
)

func singleNumbers(nums []int) []int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	index := 0
	// 循环直到最右边为1为止，此时index为从右往左第一个为 1 的位数， 即两个数字首个不同的 bit 位
	for res&1 == 0 {
		index++
		res = res >> 1
	}
	n1, n2 := 0, 0
	for i := 0; i < len(nums); i++ {
		// 如果不同的 bit 位是0
		if (nums[i]>>index)&1 == 0 {
			n1 ^= nums[i]
		} else {
			n2 ^= nums[i]
		}
	}
	return []int{n1, n2}
}

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))              // [1,6]
	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3})) // [2,10]
}
