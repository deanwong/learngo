package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	pre := dummy
	cur := head
	for cur != nil {
		next := cur.Next
		if val == cur.Val {
			pre.Next = next
		}
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

func transfrom(nums []int) *ListNode {
	dummy := &ListNode{Val: 0}
	n := len(nums)
	if n == 0 {
		return &ListNode{Val: 0}
	}
	cur := dummy
	for i := 0; i < n; i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

func (l *ListNode) String() string {
	nums := make([]int, 0)
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return fmt.Sprintf("%v", nums)
}

func main() {
	fmt.Println(deleteNode(transfrom([]int{4, 5, 1, 9}), 5)) // [4 1 9]
}
