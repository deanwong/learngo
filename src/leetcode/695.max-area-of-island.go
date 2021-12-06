package main

import "fmt"

// O(n*m)
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 {
				var area int
				dfs(grid, x, y, &area)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

func dfs(grid [][]int, x int, y int, area *int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
		return
	}
	if grid[x][y] == 1 {
		// mark
		*area += 1
		grid[x][y]++
		dfs(grid, x-1, y, area)
		dfs(grid, x+1, y, area)
		dfs(grid, x, y-1, area)
		dfs(grid, x, y+1, area)
	}
	return
}

func main() {
	// 6
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}
