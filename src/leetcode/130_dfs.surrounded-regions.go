package main

import (
	"fmt"
)

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(board, i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(board, m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x int, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[x]) {
		return
	}
	if board[x][y] != 'O' {
		return
	}
	board[x][y] = '-'
	dfs(board, x-1, y)
	dfs(board, x+1, y)
	dfs(board, x, y-1)
	dfs(board, x, y+1)
}

func solve2(board [][]byte) {
	m, n := len(board), len(board[0])
	queue := make([][]int, 0)
	var mark byte
	mark = '-'
	// 边界的 O 加入队列，标记为 -
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
			board[i][0] = mark
		}
		if board[i][n-1] == 'O' {
			queue = append(queue, []int{i, n - 1})
			board[i][n-1] = mark
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			queue = append(queue, []int{0, j})
			board[0][j] = mark
		}
		if board[m-1][j] == 'O' {
			queue = append(queue, []int{m - 1, j})
			board[m-1][j] = mark
		}
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		for i := 0; i < 4; i++ {
			newX, newY := x+directions[i][0], y+directions[i][1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && board[newX][newY] == 'O' {
				board[newX][newY] = mark
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {
	a := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(a)
	fmt.Println(a) // [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
}
