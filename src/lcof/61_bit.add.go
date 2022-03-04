package main

import (
	"fmt"
)

func add(a int, b int) int {
	// 硬背
	for b != 0 {
		// 进位和，与运算并左移
		c := a & b << 1
		// 非进位和，异或运算
		n := a ^ b
		a = n
		b = c
	}
	return a
}

func main() {
	fmt.Println(add(1, 1))   // 2
	fmt.Println(add(20, 17)) // 2
}
