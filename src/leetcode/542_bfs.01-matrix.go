package main

import (
	"fmt"
)

func updateMatrix(mat [][]int) [][]int {
	queue := make([][]int, 0)
	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] == 0 {
				queue = append(queue, []int{x, y})
			} else {
				// not visited
				mat[x][y] = -1
			}
		}
	}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		if newX := x - 1; newX >= 0 && mat[newX][y] == -1 {
			mat[newX][y] = mat[x][y] + 1
			queue = append(queue, []int{newX, y})
		}
		if newX := x + 1; newX < len(mat) && mat[newX][y] == -1 {
			mat[newX][y] = mat[x][y] + 1
			queue = append(queue, []int{newX, y})
		}
		if newY := y - 1; newY >= 0 && mat[x][newY] == -1 {
			mat[x][newY] = mat[x][y] + 1
			queue = append(queue, []int{x, newY})
		}
		if newY := y + 1; newY < len(mat[x]) && mat[x][newY] == -1 {
			mat[x][newY] = mat[x][y] + 1
			queue = append(queue, []int{x, newY})
		}
	}

	return mat
}

func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 1}, {1, 1, 0}}))
}
