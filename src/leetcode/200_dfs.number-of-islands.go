package main

import "fmt"

// O(n*m)
func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, x int, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
		return
	}
	if grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}

func numIslands2(grid [][]byte) int {
	return bfs(grid)
}

func bfs(grid [][]byte) int {
	queue := make([][]int, 0)
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				queue = append(queue, []int{i, j})
				// visited
				grid[i][j] = '2'
				ans++
				for len(queue) > 0 {
					point := queue[0]
					queue = queue[1:]
					x, y := point[0], point[1]
					if newX := x - 1; newX >= 0 && grid[newX][y] == '1' {
						queue = append(queue, []int{newX, y})
						grid[newX][y] = '2'
					}
					if newX := x + 1; newX < len(grid) && grid[newX][y] == '1' {
						queue = append(queue, []int{newX, y})
						grid[newX][y] = '2'
					}
					if newY := y - 1; newY >= 0 && grid[x][newY] == '1' {
						queue = append(queue, []int{x, newY})
						grid[x][newY] = '2'
					}
					if newY := y + 1; newY < len(grid[0]) && grid[x][newY] == '1' {
						queue = append(queue, []int{x, newY})
						grid[x][newY] = '2'
					}
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numIslands2([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}})) // 1
	fmt.Println(numIslands2([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})) // 3
}
