package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0], dp[i][i] = 1, 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp[rowIndex]
}

func getRow_plus(rowIndex int) []int {
	var pre, cur []int
	for i := 0; i < rowIndex+1; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return pre
}

func getRow_plus2(rowIndex int) []int {
	dp := make([]int, rowIndex+1)
	dp[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			dp[j] = dp[j-1] + dp[j]
		}
		fmt.Println(dp)
	}
	return dp
}

func main() {
	fmt.Println(getRow_plus2(3)) // [1,3,3,1]
	fmt.Println(getRow(0))       // [1]
	fmt.Println(getRow(1))       // [1,1]
}
