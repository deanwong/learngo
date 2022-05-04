package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	l, r, t, b := 0, n-1, 0, m-1
	x := 0
	ans := make([]int, m*n)
	for {
		for i := l; i <= r; i++ {
			ans[x] = matrix[t][i]
			x++
		}
		t++
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			ans[x] = matrix[i][r]
			x++
		}
		r--
		if l > r {
			break
		}
		for i := r; i >= l; i-- {
			ans[x] = matrix[b][i]
			x++
		}
		b--
		if t > b {
			break
		}
		for i := b; i >= t; i-- {
			ans[x] = matrix[i][l]
			x++
		}
		l++
		if l > r {
			break
		}

	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
