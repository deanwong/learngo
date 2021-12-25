package main

import "fmt"

// O(n * 2^n) 其中 n 为图中点的数量
func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
	n := len(graph)
	if n <= 0 {
		return ans
	}
	var backtrace func(from int, path []int)
	backtrace = func(from int, path []int) {
		if from == n-1 {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			ans = append(ans, pathCopy)
			return
		}
		// 城市的数量
		for i := 0; i < len(graph[from]); i++ {
			to := graph[from][i]
			path = append(path, to)
			backtrace(to, path)
			path = path[:len(path)-1]
		}
	}
	backtrace(0, []int{0})
	return ans
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))               // [[0 1 3] [0 2 3]]
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})) // [[0 1] [0 3 4] [0 1 4 3] [0 1 2 3 4] [0 1 4]]
	fmt.Println(allPathsSourceTarget([][]int{{1}, {}}))                            // [[0 1]]
	fmt.Println(allPathsSourceTarget([][]int{{1, 2, 3}, {2}, {3}, {}}))            // [[0 1 2 3] [0 2 3] [0 3]]
}
