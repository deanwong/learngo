package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归的函数要干什么？
 - 函数的作用是判断传入的树是否平衡
 - 输入：TreeNode
 - 输出：是：true，不是：false
递归停止的条件是什么？
 - 节点为nil -> true
 - 节点左子树和右子树的高度差大于 1 -> false
 - 点左子树和右子树的高度差小于 1 -> 继续递归左子树 && 递归右子树
从某层到下一层的关系是什么？
 - 判断左右子树是否平衡
*/
func isBalanced(root *TreeNode) bool {
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	if root == nil {
		return true
	}
	return math.Abs(float64(maxDepth(root.Left)-maxDepth(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced2(root *TreeNode) bool {
	/*
		递归的函数要干什么？
		- 函数的作用是判断传入的树是否平衡
		- 输入：TreeNode
		- 输出：如果平衡返回深度，不平衡返回-1
		递归停止的条件是什么？
		- 节点为nil -> 0
		- 自底向上计算
		- 节点左子树和右子树的高度差大于 1 -> -1
		- 节点左子树和右子树的高度差小于 1 -> 返回子高度
	*/
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		if left == -1 {
			return -1
		}
		right := dfs(root.Right)
		if right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) <= 1 {
			return max(left, right) + 1
		}
		return -1
	}
	return dfs(root) != -1
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
	/*
	    3
	   / \
	  9  20
	    /  \
	   15   7
	*/
	// var root *TreeNode
	// fmt.Println(isBalanced(array2binarytree([]int{3, 9, 20, -1, -1, 15, 7}, 0, root))) // true
	// fmt.Println(isBalanced2(array2binarytree([]int{3, 9, 20, -1, -1, 15, 7}, 0, root))) // true
	/*
			    1
		      / \
		     2   2
		    / \
		   3   3
		  / \
		 4   4
	*/
	var root2 *TreeNode
	fmt.Println(isBalanced(array2binarytree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, 0, root2)))  // false
	fmt.Println(isBalanced2(array2binarytree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, 0, root2))) // false
}
