package main

import "fmt"

func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		ans++
		// 以下操作相当于让num最右的 1 变成 0
		num = num & (num - 1)
	}
	return ans
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011)) // 3
	fmt.Println(hammingWeight(00000000000000000000000010000000)) // 1
}
