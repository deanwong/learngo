package main

import "fmt"

// O(n!*n) n!表示递归次数，n表示复制答案的次数
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	dfs(&ans, nums, len(nums), []int{})
	return ans
}

func dfs(ans *[][]int, nums []int, n int, permutation []int) {
	if n == 0 {
		val := make([]int, len(permutation))
		copy(val, permutation)
		*ans = append(*ans, val)
		return
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == -99 {
			continue
		}
		permutation = append(permutation, nums[i])
		val := make([]int, len(nums))
		copy(val, nums)
		val[i] = -99
		dfs(ans, val, n-1, permutation)
		permutation = permutation[:len(permutation)-1]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
