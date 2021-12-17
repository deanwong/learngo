package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(r*s)
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

func array2binarytree(array []int, idx int, node *TreeNode) *TreeNode {
	if idx < len(array) {
		if array[idx] == 0 {
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
	var subRoot *TreeNode
	fmt.Print(isSubtree(array2binarytree([]int{3, 4, 5, 1, 2}, 0, root), array2binarytree([]int{4, 1, 2}, 0, subRoot))) // [1,#,2,3,#,4,5,7,#]
}
