package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	layerIndex := 0
	for len(queue) > 0 {
		curLayer := queue
		queue = nil
		layerIndex++
		curLayerArray := make([]int, len(curLayer))
		for i := 0; i < len(curLayer); i++ {
			cur := curLayer[i]
			if layerIndex%2 == 0 {
				curLayerArray[len(curLayer)-1-i] = cur.Val
			} else {
				curLayerArray[i] = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, curLayerArray)
	}
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
	fmt.Println(levelOrder(array2binarytree([]int{3, 9, 20, -1, -1, 15, 7}, 0, root))) // [[3] [20 9] [15 7]]
}
