package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// O(n)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLayer := queue
		queue = nil
		for i := 0; i < len(curLayer); i++ {
			if i < len(curLayer)-1 {
				curLayer[i].Next = curLayer[i+1]
			}
			if curLayer[i].Left != nil {
				queue = append(queue, curLayer[i].Left)
			}
			if curLayer[i].Right != nil {
				queue = append(queue, curLayer[i].Right)
			}
		}
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	// loop from left in every layer
	for left := root; left.Left != nil; left = left.Left {
		for node := left; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

func array2binarytree(array []int, idx int, node *Node) *Node {
	if idx < len(array) {
		node = &Node{Val: array[idx]}
		node.Left = array2binarytree(array, 2*idx+1, node.Left)
		node.Right = array2binarytree(array, 2*idx+2, node.Right)
	}
	return node
}

func (n *Node) String() string {
	nums := make([]int, 0)
	quene := make([]*Node, 0)
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
	array := []int{1, 2, 3, 4, 5, 6, 7}
	var root *Node
	fmt.Print(connect2(array2binarytree(array, 0, root)))
}
