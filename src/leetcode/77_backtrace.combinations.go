package main

import "fmt"

// O((n_k)*k) n_k 表示从n个元素里面取k个元素组合，k为复制答案的次数
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	dfs(&ans, n, k, 1, []int{})
	return ans
}

func dfs(ans *[][]int, max int, k int, idx int, combine []int) {
	if k == 0 {
		val := make([]int, len(combine))
		copy(val, combine)
		*ans = append(*ans, val)
		return
	}
	for ; idx <= max-k+1; idx++ {
		combine = append(combine, idx)
		dfs(ans, max, k-1, idx+1, combine)
		combine = combine[:len(combine)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}
