package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归的函数要干什么？
函数的作用是判断传入的两个树是否镜像。
输入：TreeNode left, TreeNode right
输出：是：true，不是：false
递归停止的条件是什么？
左节点和右节点都为空 -> 倒底了都长得一样 ->true
左节点为空的时候右节点不为空，或反之 -> 长得不一样-> false
左右节点值不相等 -> 长得不一样 -> false
从某层到下一层的关系是什么？
要想两棵树镜像，那么一棵树左边的左边要和二棵树右边的右边镜像，一棵树左边的右边要和二棵树右边的左边镜像
调用递归函数传入左左和右右
调用递归函数传入左右和右左
只有左左和右右镜像且左右和右左镜像的时候，我们才能说这两棵树是镜像的
调用递归函数，我们想知道它的左右孩子是否镜像，传入的值是root的左孩子和右孩子。这之前记得判个root==null。
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(a *TreeNode, b *TreeNode) bool
	helper = func(a *TreeNode, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		return helper(a.Left, b.Right) && helper(a.Right, b.Left)
	}
	return helper(root.Left, root.Right)
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
	var a *TreeNode
	fmt.Println(isSymmetric(array2binarytree([]int{1, 2, 2, 3, 4, 4, 3}, 0, a))) // true
	var b *TreeNode
	fmt.Println(isSymmetric(array2binarytree([]int{1, 2, 2, -1, 3, -1, 3}, 0, b))) // false
}
