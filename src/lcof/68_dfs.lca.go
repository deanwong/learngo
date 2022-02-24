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
	// LCA 也包含自己
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}
	// 去该节点的左子树上找
	left := lowestCommonAncestor(root.Left, p, q)
	// 去该节点的右子树上找
	right := lowestCommonAncestor(root.Right, p, q)
	// 左右均有，说明该节点即为最近公共祖先，因为p和q都是它的子树
	if left != nil && right != nil {
		return root
	}
	// 左子树上没有，说明在右子树上
	if left == nil {
		return right
	}
	// 右子树上没有，说明在左子树上
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
	node := lowestCommonAncestor(array2binarytree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 0, root), &TreeNode{Val: 5}, &TreeNode{Val: 1})
	fmt.Println(node.Val) // 3
	node2 := lowestCommonAncestor(array2binarytree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 0, root), &TreeNode{Val: 5}, &TreeNode{Val: 4})
	fmt.Println(node2.Val) // 5
}
