package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历结果的第一位肯定是根
	root := &TreeNode{preorder[0], nil, nil}
	// 关键，寻找中序遍历结果中根的位置，然后根据跟节点的位置再到前序遍历中确定根的左右子树节点的数量
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	// 左子节点
	left := inorder[:i]
	// 右子节点
	right := inorder[i+1:]
	// 递归建立左子树，前序遍历取 第二个 直到 left 个
	root.Left = buildTree(preorder[1:len(left)+1], left)
	// 递归建立右子树，前序遍历取 left+1 直到最后
	root.Right = buildTree(preorder[len(left)+1:], right)
	return root
}

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	fmt.Println(buildTree([]int{-1}, []int{-1}))
}
