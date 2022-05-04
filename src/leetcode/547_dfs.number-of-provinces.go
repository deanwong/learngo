package main

import "fmt"

// O(n*m)
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		// not visit
		if visited[i] != 1 {
			dfs(isConnected, visited, i)
			ans++
		}
		// fmt.Println(visited)
	}
	return ans
}

func dfs(isConnected [][]int, visited []int, from int) {
	visited[from] = 1
	for to := 0; to < len(isConnected[from]); to++ {
		if isConnected[from][to] == 1 && visited[to] != 1 {
			dfs(isConnected, visited, to)
		}
	}
}

// BFS
func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected)
	ans := 0
	visited := make([]int, n)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if visited[i] != 1 {
			ans++
			queue = append(queue, i)
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				visited[from] = 1
				for to := 0; to < len(isConnected[from]); to++ {
					if isConnected[from][to] == 1 && visited[to] != 1 {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return ans
}

// UnionFind
func findCircleNum3(isConnected [][]int) int {
	ans := 0
	n := len(isConnected)
	// 类似map，下标就是key， 值就是value
	parent := make([]int, n)
	find := func(x int) int {
		// 如果找不到自己就是自己的顶点
		root := x
		// 先找到顶点
		for parent[root] != -1 {
			root = parent[root]
		}
		for x != root {
			originParent := parent[x]
			// 压缩路径，让节点的顶点直接为根顶点
			parent[x] = root
			x = originParent
		}
		return root
	}
	union := func(from, to int) {
		fromRoot := find(from)
		toRoot := find(to)
		if fromRoot != toRoot {
			parent[fromRoot] = toRoot
			ans--
		}
	}
	add := func(x int) {
		// dummy顶点为-1
		parent[x] = -1
		ans++
	}
	for from := 0; from < n; from++ {
		add(from)
		// 特别注意 to < from
		for to := 0; to < from; to++ {
			if isConnected[from][to] == 1 {
				union(from, to)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findCircleNum3([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})) // 2
	fmt.Println(findCircleNum3([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // 3
}
