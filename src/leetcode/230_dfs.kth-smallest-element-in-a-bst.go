package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]int, 0)
	dfs(root, &stack, k)
	return stack[len(stack)-1]
}

func dfs(root *TreeNode, stack *[]int, k int) {
	if root == nil {
		return
	}
	dfs(root.Left, stack, k)
	if k == len(*stack) {
		return
	}
	*stack = append(*stack, root.Val)
	dfs(root.Right, stack, k)
}

func array2binarytree(array []int, idx int, node *TreeNode) *TreeNode {
	if idx < len(array) {
		if array[idx] == -1 {
			return nil
		}
		node = &TreeNode{Val: array[idx]}
		node.Left = array2binarytree(array, 2*idx+1, node.Left)
		node.Right = array2binarytree(array, 2*idx+2, node.Right)
	}
	return node
}

func main() {
	var root1 *TreeNode
	fmt.Println(kthSmallest(array2binarytree([]int{3, 1, 4, -1, 2}, 0, root1), 1)) //1
	var root2 *TreeNode
	fmt.Println(kthSmallest(array2binarytree([]int{5, 3, 6, 2, 4, -1, -1, 1}, 0, root2), 3)) //3
}
