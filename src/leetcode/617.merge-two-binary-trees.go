package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	node := &TreeNode{root1.Val + root2.Val, nil, nil}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)
	return node
}

func array2binarytree(array []int, idx int, node *TreeNode) *TreeNode {
	if idx < len(array) {
		node = &TreeNode{array[idx], nil, nil}
		node.Left = array2binarytree(array, 2*idx+1, node.Left)
		node.Right = array2binarytree(array, 2*idx+2, node.Right)
	}
	return node
}

func (n *TreeNode) String() string {
	nums := make([]int, 0)
	quene := make([]*TreeNode, 0)
	quene = append(quene, n)
	for len(quene) > 0 {
		cur := quene[0]
		quene = quene[1:]
		nums = append(nums, cur.Val)
		if cur.Left != nil {
			quene = append(quene, cur.Left)
		}
		if cur.Right != nil {
			quene = append(quene, cur.Right)
		}
	}
	return fmt.Sprintf("%v", nums)
}

func main() {
	array1 := []int{1, 3, 2, 5}
	var root1 *TreeNode
	array2 := []int{2, 1, 3, 0, 4, 0, 7}
	var root2 *TreeNode
	fmt.Print(mergeTrees(array2binarytree(array1, 0, root1), array2binarytree(array2, 0, root2)))
}
