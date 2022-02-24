package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 当 p, q 都在 root的 右子树 中，则开启递归 root.right 并返回；
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// 当 p, q 都在 root 的 左子树 中，则开启递归 root.left 并返回；
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	// 恰好在中间就是LCA，因为BST特性 左 < 根 < 右
	return root
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
	node := lowestCommonAncestor(array2binarytree([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 0, root), &TreeNode{Val: 2}, &TreeNode{Val: 8})
	fmt.Println(node.Val) // 6
	node2 := lowestCommonAncestor(array2binarytree([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, 0, root), &TreeNode{Val: 2}, &TreeNode{Val: 4})
	fmt.Println(node2.Val) // 2
}
