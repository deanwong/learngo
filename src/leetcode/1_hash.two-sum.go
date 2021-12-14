package main

import "fmt"

// O(n)
func twosum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if v, ok := dict[target-num]; ok {
			return []int{v, i}
		}
		dict[num] = i
	}
	return nil
}

func main() {

	fmt.Println(twosum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twosum([]int{3, 2, 4}, 6))
	fmt.Println(twosum([]int{3, 3}, 6))
}
