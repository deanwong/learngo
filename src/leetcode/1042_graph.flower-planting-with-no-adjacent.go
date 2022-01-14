package main

import "fmt"

// O(m*n)
func gardenNoAdj(n int, paths [][]int) []int {
	// 做邻接表 graph[i] 表示 i 花园关联的位置，首位留空
	graph := make([][]int, n+1)
	for i := 0; i < len(paths); i++ {
		graph[paths[i][0]] = append(graph[paths[i][0]], paths[i][1])
		graph[paths[i][1]] = append(graph[paths[i][1]], paths[i][0])
	}
	// fmt.Println(graph)
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 用下标表示是否使用，如 cadidates[1] > 0 表示 1 号花被使用了
		cadidates := make([]int, 5)
		// 遍历 i 号花园的相邻花园
		for j := 0; j < len(graph[i]); j++ {
			// 如果相邻的花园有使用过的花，那么标记
			used := ans[graph[i][j]]
			if used != 0 {
				cadidates[used]++
			}
		}
		// fmt.Printf("ans %v\n", ans)
		// fmt.Printf("cadidates %v\n", cadidates)
		for k := 1; k < len(cadidates); k++ {
			if cadidates[k] == 0 {
				ans[i] = k
				break
			}
		}
	}
	return ans[1:]
}

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))                         // [1,2,3]
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {3, 4}}))                                 // [1,2,1,1]
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}})) // [1,2,3,4]
}
