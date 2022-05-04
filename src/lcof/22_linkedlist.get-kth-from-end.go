package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	a, b := dummy, dummy
	for i := 1; i <= k && b != nil; i++ {
		b = b.Next
	}
	for b != nil {
		a = a.Next
		b = b.Next
	}
	return a
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
	fmt.Println(getKthFromEnd(transfrom([]int{1, 2, 3, 4, 5}), 2)) // [4 5]
}
