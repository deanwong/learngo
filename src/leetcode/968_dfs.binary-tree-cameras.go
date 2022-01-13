package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total int

func minCameraCover(root *TreeNode) int {
	total = 0
	// 若root的状态是0，即未被监控，那么需要安装一个摄像头，计数器+1
	if dfs(root) == 0 {
		total++
	}
	return total
}

/**
 * 每个节点三种状态：
 * 0-->未被监控
 * 1-->已被监控（该节点未安装摄像头）
 * 2-->已被监控（该节点已安装摄像头，计数器+1）
 */
func dfs(root *TreeNode) int {
	if root == nil {
		return 1
	}
	l := dfs(root.Left)
	r := dfs(root.Right)

	// 如果左子节点、右子节点的和大于3表示其中一个一定是2已经安装了摄像头，返回状态1
	if l+r >= 3 {
		return 1
	}
	// 如果左子结点、右子节点的状态都是1已被监控（未安装摄像头）
	// 那么它们的父节点（当前节点）的状态为未被监控，直接返回状态0
	if l == 1 && r == 1 {
		return 0
	} else {
		// 左子结点和右子节点至少一个未被监控的情况，当前节点需要安装摄像头，计数器+1，返回状态2
		total++
		return 2
	}
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
	fmt.Println(minCameraCover(array2binarytree([]int{0, 0, -1, 0, 0}, 0, root1))) //1
	var root2 *TreeNode
	fmt.Println(minCameraCover(array2binarytree([]int{0, 0, -1, 0, -1, 0, -1, -1, 0}, 0, root2))) //2
}
