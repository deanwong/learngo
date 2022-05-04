package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	ans := make([][]int, 0)
	var backtrace func(root *TreeNode, temp []int, sum int)
	backtrace = func(root *TreeNode, temp []int, sum int) {
		if root.Left == nil && root.Right == nil && sum == target {
			tempCopy := make([]int, len(temp))
			copy(tempCopy, temp)
			ans = append(ans, tempCopy)
			return
		}
		if root.Left != nil {
			temp = append(temp, root.Left.Val)
			backtrace(root.Left, temp, sum+root.Left.Val)
			temp = temp[:len(temp)-1]
		}
		if root.Right != nil {
			temp = append(temp, root.Right.Val)
			backtrace(root.Right, temp, sum+root.Right.Val)
			temp = temp[:len(temp)-1]
		}
	}
	if root == nil {
		return [][]int{}
	}
	backtrace(root, []int{root.Val}, root.Val)
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
	fmt.Println(pathSum(array2binarytree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1}, 0, root), 22)) // [[5,4,11,2],[5,8,4,5]]
}
