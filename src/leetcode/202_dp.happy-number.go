package main

import "fmt"

// O(
func isHappy(n int) bool {
	num := n
	lookup := make(map[int]bool)
	for num != 1 {
		sum := 0
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num = num / 10
		}
		if _, ok := lookup[sum]; ok {
			return false
		} else {
			lookup[sum] = true
		}
		num = sum
	}
	return num == 1
}

func main() {
	fmt.Println(isHappy(19)) // true
	fmt.Println(isHappy(2))  // false
}
