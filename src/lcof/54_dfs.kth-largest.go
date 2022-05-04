package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode, stack *[]int)
	dfs = func(root *TreeNode, stack *[]int) {
		if root == nil {
			return
		}
		dfs(root.Right, stack)
		if len(*stack) == k {
			return
		}
		*stack = append(*stack, root.Val)
		dfs(root.Left, stack)
	}
	stack := make([]int, 0)
	dfs(root, &stack)
	return stack[len(stack)-1]
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
	var root *TreeNode
	fmt.Println(kthLargest(array2binarytree([]int{3, 1, 4, -1, 2, -1, -1}, 0, root), 1)) // 4
	var root2 *TreeNode
	fmt.Println(kthLargest(array2binarytree([]int{5, 3, 6, 2, 4, -1, -1, 1}, 0, root2), 3)) // 4
}
