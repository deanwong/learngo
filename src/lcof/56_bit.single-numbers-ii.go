package main

import (
	"fmt"
)

func singleNumbers2(nums []int) int {
	// 可以设计一种逻辑，使数字出现 3 次时，该逻辑的结果为 0（即只有 0，1，2 三种状态）
	// 其实就是一个 三进制
	// 一位二进制数只能存储 0 和 1 两种状态，所以我们需要用到两位二进制
	// 设两位二进制数的高位为 A，低位为 B。C 是输入变量
	// 表示的三种情况为 ： 0次：00(A=0,B=0), 1次：01(A=0,B=1), 2次：10(A=1,B=0)
	// 注：11(A=1,B=1) 为无效输入

	// 画出关于 A 的卡诺图（AB为11的结果是不重要的，用 x 表示）：
	//  AB\C |  0  |  1
	//  =================
	//    00 |  0  |  0
	//    01 |  0  |  1        ====> 得到 F = ABC + A'BC + ABC' + AB'C' = BC + AC'
	//    11 |  1  |  1
	//    10 |  1  |  0

	//  画出关于 B 的卡诺图
	//  AB\C |  0  |  1
	//  =================
	//    00 |  0  |  1
	//    01 |  1  |  0        ====> 得到 F = ABC' + A'BC' + A'B'C = BC' + A'B'C
	//    11 |  1  |  0
	//    10 |  0  |  0

	// 很明显啊，我们需要的就是只出现一次的情况 01（A=0，B=1），即 B 的结果
	A, B := 0, 0
	for i := 0; i < len(nums); i++ {
		C := nums[i]
		temp := A
		A = (B & C) | (A & ^C)
		B = (B & ^C) | (^temp & ^B & C)
	}
	return B
}

func singleNumbers(nums []int) int {
	lookup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		lookup[num]++
		if lookup[num] == 3 {
			delete(lookup, num)
		}
	}
	for k, _ := range lookup {
		return k
	}
	return 0
}

func main() {
	fmt.Println(singleNumbers([]int{3, 4, 3, 3}))          // [4]
	fmt.Println(singleNumbers([]int{9, 1, 7, 9, 7, 9, 7})) // [1]
}
