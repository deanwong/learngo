package main

import "fmt"

// O(1)
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}
