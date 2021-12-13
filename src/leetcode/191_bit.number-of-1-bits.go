package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if num&(1<<i) > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}
