package main

import "fmt"

func movingCount(m int, n int, k int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	sum := func(x, y int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x = x / 10
		}
		for y != 0 {
			res += y % 10
			y = y / 10
		}
		return res
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || y < 0 || x >= m || y >= n || visited[x][y] || sum(x, y) > k {
			return 0
		}
		visited[x][y] = true
		res := 0
		for i := 0; i < 4; i++ {
			newX, newY := x+directions[i][0], y+directions[i][1]
			res += dfs(newX, newY)
		}
		return res + 1
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(movingCount(2, 3, 1)) // 3
	fmt.Println(movingCount(3, 1, 0)) // 1
}
