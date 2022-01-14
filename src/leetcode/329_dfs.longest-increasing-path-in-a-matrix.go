package main

import (
	"fmt"
)

func longestIncreasingPath(matrix [][]int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(matrix), len(matrix[0])
	// 计算以每个节点开头的递增序列的长度
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(x int, y int) int
	dfs = func(x, y int) int {
		// fmt.Printf("before x %v y %v visited %v\n", x, y, memo)
		if memo[x][y] > 0 {
			return memo[x][y]
		}
		// 长度至少为1
		maxLen := 1
		for i := 0; i < 4; i++ {
			newX, newY := x+directions[i][0], y+directions[i][1]
			// 寻找比自己大的周围并递归，在返回值上+1
			if newX >= 0 && newX < m && newY >= 0 && newY < n && matrix[newX][newY] > matrix[x][y] {
				maxLen = max(maxLen, dfs(newX, newY)+1)
			}
		}
		memo[x][y] = maxLen
		// fmt.Printf("after x %v y %v visited %v\n", x, y, memo)
		return memo[x][y]
	}
	maxLen := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxLen = max(maxLen, dfs(i, j))
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})) // 4 [1, 2, 6, 9]
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})) // 4 [3, 4, 5, 6]
	fmt.Println(longestIncreasingPath([][]int{{1}}))                             // 1
}
