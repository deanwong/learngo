package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
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
	fmt.Println(hasPathSum(array2binarytree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, 0, root1), 22)) //true
	var root2 *TreeNode
	fmt.Println(hasPathSum(array2binarytree([]int{1, 2, 3}, 0, root2), 5)) // false
	var root3 *TreeNode
	fmt.Println(hasPathSum(array2binarytree([]int{}, 0, root3), 0)) // false
	var root4 *TreeNode
	fmt.Println(hasPathSum(array2binarytree([]int{1, 2}, 0, root4), 1)) // false
}
