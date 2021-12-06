package main

import "fmt"

// O(n*m)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		dfs(image, sr, sc, color, newColor)
	}
	return image
}

func dfs(image [][]int, x int, y int, color int, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[x]) {
		return
	}
	if image[x][y] == color {
		image[x][y] = newColor
		dfs(image, x-1, y, color, newColor)
		dfs(image, x+1, y, color, newColor)
		dfs(image, x, y-1, color, newColor)
		dfs(image, x, y+1, color, newColor)
	}

}

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2))
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
}
