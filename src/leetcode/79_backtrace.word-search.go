package main

import "fmt"

// 一个非常宽松的上界为 O(M * N * 3^L) 其中 M, N 为网格的长度与宽度，L 为字符串的长度。在每次调用函数 backtrace 时，除了第一次可以进入 4 个分支以外，其余时间我们最多会进入 3 个分支（因为每个位置只能使用一次，所以走过来的分支没法走回去）
func exist(board [][]byte, word string) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var backtrace func(x int, y int, idx int) bool
	backtrace = func(x, y, idx int) bool {
		if board[x][y] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		visited[x][y] = true
		fmt.Printf("before  %v y %v visited %v\n", x, y, visited)
		for i := 0; i < 4; i++ {
			newX, newY := directions[i][0]+x, directions[i][1]+y
			if newX >= 0 && newX < m && newY >= 0 && newY < n && !visited[newX][newY] {
				if backtrace(newX, newY, idx+1) {
					return true
				}
			}
		}
		visited[x][y] = false
		fmt.Printf("after x %v y %v visited %v\n", x, y, visited)
		return false
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
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))    // true
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))   // fasle
	fmt.Println(exist([][]byte{{'a', 'a'}}, "aa"))                                                           // true
}

func exist2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	queue := make([][]int, 0)
	level := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[level] {
				// 有bug，从level1开始需要个各自的访问类别 visited
				board[i][j] = '-'
				queue = append(queue, []int{i, j})
			}
		}
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		level++
		curLayer := queue
		queue = nil
		for i := 0; i < len(curLayer); i++ {
			point := curLayer[i]
			x, y := point[0], point[1]
			for j := 0; j < 4; j++ {
				newX, newY := x+directions[j][0], y+directions[j][1]
				fmt.Printf("newX:%v,newY:%v\n", newX, newY)
				if newX >= 0 && newX < m && newY >= 0 && newY < n && level < len(word) && board[newX][newY] == word[level] {
					fmt.Printf("newX:%v,newY:%v, word %v, level %v\n", newX, newY, string(word[level]), level)
					board[newX][newY] = '-'
					queue = append(queue, []int{newX, newY})
				}
			}
		}
	}
	return level == len(word)
}
