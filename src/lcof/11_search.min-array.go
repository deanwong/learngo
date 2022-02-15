package main

import "fmt"

func minArray(numbers []int) int {
	n := len(numbers)
	if n == 0 {
		return 0
	}
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if numbers[mid] < numbers[hi] {
			// mid 的右边一定不是最小数字，mid 有可能是，下一轮搜索区间是 [lo, mid]
			hi = mid
		} else if numbers[mid] > numbers[hi] {
			// 最小数字肯定在右边，抛弃左边
			lo = mid + 1
		} else {
			// 此时 numbers[mid] == numbers[hi]， 只能排除hi
			hi--
		}
	}
	return numbers[lo]
}

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 6, 7, 1, 2})) // 1
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))       // 0
	fmt.Println(minArray([]int{1, 3, 5}))             // 1
	fmt.Println(minArray([]int{10, 1, 10, 10, 10}))   // 1
}
