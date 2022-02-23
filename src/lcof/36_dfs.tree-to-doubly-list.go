package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	pre := dummy
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		root.Left = pre
		pre.Right = root
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	// 连接首尾
	head := dummy.Right
	head.Left = pre
	pre.Right = head
	return head
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
	fmt.Println(treeToDoublyList(array2binarytree([]int{4, 2, 6, 1, 3, 5, 7}, 0, root)))
}
