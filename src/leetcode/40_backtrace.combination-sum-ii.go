package main

import (
	"fmt"
	"sort"
)

// O(n * 2^n)
func combinationSum2(candidates []int, target int) [][]int {
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
		// dfs的start必须是之后的数字，进入dfs的条件是当前数字小于剩余的target
		for i := start; i < n && candidates[i] <= targetLeft; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			combination = append(combination, candidates[i])
			backtrace(targetLeft-candidates[i], i+1, combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtrace(target, 0, []int{})
	return ans
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)) // [[1 1 6] [1 2 5] [1 7] [2 6]]
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))        // [[1 2 2] [5]]
}
