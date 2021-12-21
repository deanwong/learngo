package main

import "fmt"

// O(n * 2^n) 其中 n 为图中点的数量
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans := make([][]int, 0)
	if n <= 0 {
		return ans
	}
	dfs(&ans, graph, n, 0, []int{0})
	return ans
}

func dfs(ans *[][]int, graph [][]int, n int, from int, path []int) {
	if from == n-1 {
		val := make([]int, len(path))
		copy(val, path)
		*ans = append(*ans, val)
		return
	}
	for i := 0; i < len(graph[from]); i++ {
		path = append(path, graph[from][i])
		dfs(ans, graph, n, graph[from][i], path)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))               // [[0 1 3] [0 2 3]]
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})) // [[0 1] [0 3 4] [0 1 4 3] [0 1 2 3 4] [0 1 4]]
	fmt.Println(allPathsSourceTarget([][]int{{1}, {}}))                            // [[0 1]]
	fmt.Println(allPathsSourceTarget([][]int{{1, 2, 3}, {2}, {3}, {}}))            // [[0 1 2 3] [0 2 3] [0 3]]
}
