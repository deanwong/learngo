package main

import "fmt"

// O(n)
func twosum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		left := target - nums[i]
		if _, ok := lookup[left]; ok {
			return []int{lookup[left], i}
		} else {
			lookup[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twosum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twosum([]int{3, 2, 4}, 6))      //[1 2]
	fmt.Println(twosum([]int{3, 3}, 6))         //[0 1]
}
