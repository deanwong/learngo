package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var isSameTree func(a, b *TreeNode) bool
	isSameTree = func(a, b *TreeNode) bool {
		if b == nil {
			return true
		}
		// b 不是 nil 但 a 是，说明深度不同
		if a == nil {
			return false
		}
		return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
	}
	if A == nil || B == nil {
		return false
	}
	return isSameTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	var isSameTree func(a, b *TreeNode) bool
	isSameTree = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		ans := a.Val == b.Val
		if b.Left != nil {
			ans = ans && isSameTree(a.Left, b.Left)
		}
		if b.Right != nil {
			ans = ans && isSameTree(a.Right, b.Right)
		}
		return ans
	}
	if A == nil {
		return false
	}
	return isSameTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
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
	var A, B *TreeNode
	fmt.Println(isSubStructure(array2binarytree([]int{1, 0, 1, -4, -3}, 0, A), array2binarytree([]int{1, -4}, 0, B))) // true
}
