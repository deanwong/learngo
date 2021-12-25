package main

import (
	"fmt"
	"math"
)

// O(n)
func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	for i := h - 1; i >= 0; i-- {
		dp[i] = make([]int, len(triangle[i]))
		// bottom -> top
		for j := 0; j < len(triangle[i]); j++ {
			// bottom value is self
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				// n row's value come from n+1 rows
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func minimumTotal_fromtop(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]
	for i := 1; i < h; i++ {
		dp[i] = make([]int, len(triangle[i]))
		// top -> bottom
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			// n row's value come from n-1 rows
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < len(dp[h-1]); i++ {
		ans = min(ans, dp[h-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumTotal_fromtop([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})) // 11
	fmt.Println(minimumTotal_fromtop([][]int{{-10}}))                                // -10
	fmt.Println(minimumTotal_fromtop([][]int{{-1}, {2, 3}, {1, -1, -3}}))            // -1
}
