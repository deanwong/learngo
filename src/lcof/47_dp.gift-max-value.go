package main

import "fmt"

func maxValue(grid [][]int) int {
	// definition: dp[i][j] 表示到达 i,j 个格子的最大价值
	// formulation : dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// initial: dp[0][0] 左上角为 grid[0][0]
	// answer: dp[m-1][n-1]
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(maxValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) // 12
}
