package main

import (
	"fmt"
)

func main() {
	val := 0
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)
	numbers := fibonacci(val)
	fmt.Println(numbers)
}

func fibonacci(num int) []int {
	if num < 2 {
		return make([]int, 0)
	}
	numbers := make([]int, num)
	numbers[0] = 1
	numbers[1] = 1
	for i := 2; i < num; i++ {
		numbers[i] = numbers[i-2] + numbers[i-1]
	}
	return numbers
}
