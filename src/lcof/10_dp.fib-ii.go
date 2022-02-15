package main

import "fmt"

func numWays(n int) int {
	const mod int = 1e9 + 7
	if n == 0 || n == 1 {
		return 1
	}
	a := 1
	b := 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (a + b) % mod
		a = b
		b = ans
	}
	return ans
}

func main() {
	fmt.Println(numWays(2)) // 1
	fmt.Println(numWays(7)) // 21
	fmt.Println(numWays(0)) // 1
}
