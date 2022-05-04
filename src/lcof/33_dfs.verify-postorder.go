package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	var dfs func(left, right int) bool
	dfs = func(left, right int) bool {
		//如果left==right，就一个节点不需要判断了，如果left>right说明没有节点，
		//也不用再看了,否则就要继续往下判断
		if left >= right {
			return true
		}
		//因为数组中最后一个值postorder[right]是根节点，这里从左往右找出第一个比根节点大的值，
		//他后面的都是根节点的右子节点（包含当前值，不包含最后一个值，因为最后一个是根节点），他前面的都是根节点的左子节点
		greater := left
		root := postorder[right]
		for postorder[greater] < root {
			greater++
		}
		// greater 到 right - 1 都是右子节点，left 到 greater-1 都是左子节点
		// 需要判断右子节点是不是都比root大
		for i := greater; i < right; i++ {
			if postorder[i] < root {
				return false
			}
		}
		return dfs(left, greater-1) && dfs(greater, right-1)
	}
	return dfs(0, len(postorder)-1)
}

func main() {
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5})) // false
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5})) // true
}
