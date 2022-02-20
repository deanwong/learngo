package main

import "fmt"

func exist(board [][]byte, word string) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var backtrace func(x, y, idx int) bool
	backtrace = func(x, y, idx int) bool {
		// 终止条件
		if idx >= len(word) {
			return true
		}
		// 终止条件，坐标范围错误或者 i j 坐标对应 board 的值不等于 k 坐标对应 word 的值
		if x < 0 || y < 0 || x >= m || y >= n || board[x][y] != word[idx] || visited[x][y] {
			return false
		}
		visited[x][y] = true
		ans := false
		for i := 0; i < 4; i++ {
			newX, newY := x+directions[i][0], y+directions[i][1]
			ans = ans || backtrace(newX, newY, idx+1)
		}
		visited[x][y] = false
		return ans
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrace(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")) // true
}
