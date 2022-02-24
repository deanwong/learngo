package main

import "fmt"

func sumNums(n int) int {
	ans := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func sumNums2(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNums2(n-1)
}

func main() {
	fmt.Println(sumNums2(3)) // 6
	fmt.Println(sumNums2(9)) // 45
}
