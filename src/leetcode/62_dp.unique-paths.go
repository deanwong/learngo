package main

import "fmt"

// O(m*n)
func uniquePaths(m int, n int) int {
	// definition: dp[i][j] 表示到i,j点有几种走法
	// formulation: dp[i][j] = dp[i][j-1] + dp[i-1][j] 要么从上面来来要么从左边来
	// initial: i == 0 || j == 0 => dp[i][j] = 1 第一行和第一列都是只有 1 种走法
	// answer:  dp[m-1][n-1]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// first line only come from left, first column only come from top
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				// come from left or top
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7)) // 28
	fmt.Println(uniquePaths(3, 2)) //3
}
