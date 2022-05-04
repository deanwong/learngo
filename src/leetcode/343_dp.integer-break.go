package main

import (
	"fmt"
)

// O(n*n)
func integerBreak(n int) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// definition: dp[i] 表示 i 拆分成两个正整数的最大乘积
	// formulation: dp[i] = max(dp[i-j] * j, i-j*j) 0<j<i, dp[i-j] * j 表示对 i-j 再次拆分
	// 例如计算 dp[4]，可以取 dp[4-1] * 1 即 dp[3] * 1 或 dp[2] * 2 或 dp[3] * 3 ，和 3*1, 2*2, 1*3 其中的最大值, 比较容易忘记还要算 (i-j)*j
	// initial: dp[0],dp[1]=0
	// answer: dp[n]
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(2))  //1
	fmt.Println(integerBreak(10)) //36
}
