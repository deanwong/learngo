package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := ""
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		ans = ans + fmt.Sprint(root.Val)
		if root.Left != nil {
			ans = ans + "("
			dfs(root.Left)
			ans = ans + ")"
		}
		if root.Left == nil && root.Right != nil {
			ans = ans + "(" + ")"
		}
		if root.Right != nil {
			ans = ans + "("
			dfs(root.Right)
			ans = ans + ")"
		}
	}
	dfs(root)
	return ans
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
	fmt.Println(tree2str(array2binarytree([]int{1, 2, 3, 4}, 0, root)))
}
