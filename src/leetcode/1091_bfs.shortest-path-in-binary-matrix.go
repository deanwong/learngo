package main

import (
	"fmt"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n <= 0 || grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	grid[0][0]++
	// eight directions
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	// 0 表示没有访问过
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		for i := 0; i < len(directions); i++ {
			newX, newY := x+directions[i][0], y+directions[i][1]
			if newX >= 0 && newX < n && newY >= 0 && newY < n && grid[newX][newY] == 0 {
				grid[newX][newY] = grid[x][y] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	if grid[n-1][n-1] == 0 {
		return -1
	}
	return grid[n-1][n-1]
}

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))                  // 2
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})) // 4
	fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}})) // -1
}
