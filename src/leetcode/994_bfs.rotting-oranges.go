package main

import (
	"fmt"
)

func orangesRotting(grid [][]int) int {
	queue := make([][]int, 0)
	good := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 2 {
				queue = append(queue, []int{x, y})
			}
			if grid[x][y] == 1 {
				good++
			}
		}
	}
	if good == 0 {
		return 0
	}
	level := 0
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for n := 0; n < len(curLevel); n++ {
			x, y := curLevel[n][0], curLevel[n][1]
			fmt.Printf("x:%v,y:%v\n", x, y)
			if newX := x - 1; newX >= 0 && grid[newX][y] == 1 {
				grid[newX][y] = 2
				good--
				queue = append(queue, []int{newX, y})
			}
			if newX := x + 1; newX < len(grid) && grid[newX][y] == 1 {
				grid[newX][y] = 2
				good--
				queue = append(queue, []int{newX, y})
			}
			if newY := y - 1; newY >= 0 && grid[x][newY] == 1 {
				grid[x][newY] = 2
				good--
				queue = append(queue, []int{x, newY})
			}
			if newY := y + 1; newY < len(grid[x]) && grid[x][newY] == 1 {
				grid[x][newY] = 2
				good--
				queue = append(queue, []int{x, newY})
			}
		}
		level++
		fmt.Printf("level:%v, good:%v\n", level, good)
	}
	if good > 0 {
		return -1
	}
	// 烂完会多加一分钟，去掉
	return level - 1
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0}}))
	fmt.Println(orangesRotting([][]int{{1}}))
	fmt.Println(orangesRotting([][]int{{2}}))
}
