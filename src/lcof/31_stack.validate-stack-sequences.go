package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	j := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for j < len(popped) && len(stack) > 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(popped) == j
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
