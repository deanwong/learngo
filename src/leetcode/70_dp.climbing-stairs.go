package main

import "fmt"

// O(n)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		// f(n) = f(n-2) + f(n-1)
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func climbStairs_plus(n int) int {
	if n <= 2 {
		return n
	}
	a := 1
	b := 2
	for i := 2; i < n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	ans := 0
	dfs(&ans, n, 0)
	return ans
}

func dfs(ans *int, n int, distance int) {
	if distance > n {
		return
	}
	if distance == n {
		*ans++
		return
	}
	for step := 1; step <= 2; step++ {
		distance += step
		// fmt.Println(distance)
		dfs(ans, n, distance)
		distance -= step
		// fmt.Println(distance)
	}
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(5))
}
