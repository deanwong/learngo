package main

import "fmt"

// O(n)
func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
