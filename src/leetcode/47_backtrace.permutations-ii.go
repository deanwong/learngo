package main

import (
	"fmt"
	"sort"
)

// O(n!*n) n!表示递归次数，n表示复制答案的次数
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	var backtrace func(candidates []int, left int, permutation []int)
	backtrace = func(candidates []int, left int, permutation []int) {
		if left == 0 {
			permutationCopy := make([]int, len(permutation))
			copy(permutationCopy, permutation)
			ans = append(ans, permutationCopy)
			return
		}
		for i := 0; i < n; i++ {
			if candidates[i] == -99 {
				continue
			}
			// 去重
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			candidatesCopy := make([]int, len(candidates))
			copy(candidatesCopy, candidates)
			permutation = append(permutation, candidatesCopy[i])
			candidatesCopy[i] = -99
			backtrace(candidatesCopy, left-1, permutation)
			permutation = permutation[:len(permutation)-1]
		}
	}
	backtrace(nums, n, []int{})
	return ans
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2})) // [[1 1 2] [1 2 1] [2 1 1]]
	fmt.Println(permuteUnique([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
