package main

import (
	"fmt"
	"sort"
)

// O(n * 2^n)
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(candidates)
	sort.Ints(candidates)
	var backtrace func(targetLeft int, start int, combination []int)
	backtrace = func(targetLeft int, start int, combination []int) {
		if targetLeft == 0 {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			ans = append(ans, combinationCopy)
			return
		}
		// dfs的start必须是自己及之后的数字，进入dfs的条件是当前数字小于剩余的target
		for i := start; i < n && candidates[i] <= targetLeft; i++ {
			combination = append(combination, candidates[i])
			backtrace(targetLeft-candidates[i], i, combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtrace(target, 0, []int{})
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // [[2 2 2 2] [2 3 3] [3 5]]
	fmt.Println(combinationSum([]int{2}, 1))          // []
}
