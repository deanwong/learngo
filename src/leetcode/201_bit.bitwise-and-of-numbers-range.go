package main

import (
	"fmt"
)

// O(logn) 算法的时间复杂度取决于 m 和 n 的二进制位数，由于 m≤n，因此时间复杂度取决于 n 的二进制位数
// 取 left 和 right 二进制的公共前缀
func rangeBitwiseAnd(left int, right int) int {
	step := 0
	// 两个数字不断右移，直到相等，即数字被缩减为它们的公共前缀，记录步数
	// 最终左移步数后就是答案
	for right > left {
		left = left >> 1
		right = right >> 1
		step++
	}
	return right << step
}

// Brian Kernighan 算法
// 算法的关键在于我们每次对 number 和 number−1 之间进行按位与运算后，number 中最右边的 1 会被抹去变成 0
func rangeBitwiseAnd2(left int, right int) int {
	n := right
	for n > left {
		n = n & (n - 1)
	}
	return n
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))          // 4
	fmt.Println(rangeBitwiseAnd(0, 0))          // 0
	fmt.Println(rangeBitwiseAnd(1, 2147483647)) // 0
}
