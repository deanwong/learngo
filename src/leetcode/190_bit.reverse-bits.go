package main

import "fmt"

// O(1）
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		// 左移并加上nums最低位
		ans = ans<<1 | num&1
		num = num >> 1
	}
	return ans
}

func main() {
	fmt.Println(reverseBits(00000000000000000000000000001011))
}
