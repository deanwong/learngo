package main

import "fmt"

// O((n_k)*k) n_k 表示从n个元素里面取k个元素组合，k为复制答案的次数
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var backtrace func(start int, left int, combination []int)
	backtrace = func(start int, left int, combination []int) {
		if left == 0 {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			ans = append(ans, combinationCopy)
			return
		}
		for i := start; i <= n; i++ {
			combination = append(combination, i)
			backtrace(i+1, left-1, combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtrace(1, k, []int{})
	return ans
}

func main() {
	fmt.Println(combine(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
	fmt.Println(combine(1, 1)) // [[1]]
}
