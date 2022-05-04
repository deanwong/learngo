package main

import "fmt"

// O(n)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// definition: dp[i] 表示爬 i 级台阶的走法
	// formulation : dp[i] = dp[i-1] + dp[i-1] 爬 i 级台阶取可以上一步只走 1 级 dp[i-1] 也可以 走 2 级 dp[i-2]
	// initial: dp[0] = 0, dp[1] = 1, dp[2] = 2 (这里与斐波那契不同,斐波那契f(2)=1)
	// answer: dp[n]
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs_plus(n int) int {
	if n <= 2 {
		return n
	}
	a := 1
	b := 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

func main() {
	fmt.Println(climbStairs_plus(2)) // 2
	fmt.Println(climbStairs_plus(5)) // 8
}
