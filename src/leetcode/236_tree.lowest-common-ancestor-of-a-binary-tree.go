package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LCA
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
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
	// node := lowestCommonAncestor(array2binarytree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 0, root), &TreeNode{Val: 5}, &TreeNode{Val: 1})
	// fmt.Println(node.Val) // 3
	node2 := lowestCommonAncestor(array2binarytree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 0, root), &TreeNode{Val: 7}, &TreeNode{Val: 4})
	fmt.Println(node2.Val) // 5
}
