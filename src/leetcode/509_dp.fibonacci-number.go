package main

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func fib_dp(n int) int {
	if n < 2 {
		return n
	}
	// definition: dp[i] 表示 i 个斐波那契数列的总和
	// formulation : dp[i] = dp[i-2] + dp[i-1]
	// initial: dp[0]=0 dp[1]=1
	// answer: dp[n]
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func fib_dp_plus(n int) int {
	if n < 2 {
		return n
	}
	a := 0
	b := 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

func main() {
	fmt.Println(fib_dp_plus(2)) // 1
	fmt.Println(fib_dp_plus(3)) // 2
	fmt.Println(fib_dp_plus(4)) // 3
	fmt.Println(fib_dp_plus(8)) // 21
}
