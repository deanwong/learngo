package main

import "fmt"

// O(m*n)
func maxIncreaseKeepingSkyline(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// 计算各行和各列的最大值
	mMax := make([]int, m)
	nMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mMax[i] = max(mMax[i], grid[i][j])
			nMax[j] = max(nMax[j], grid[i][j])
		}
	}
	// fmt.Println(mMax)
	// fmt.Println(nMax)
	// 计算差值
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += min(mMax[i], nMax[j]) - grid[i][j]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}})) // 35
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}))                        // 0
}
