package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return dummy.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
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
	fmt.Println(mergeTwoLists(transfrom([]int{1, 2, 4}), transfrom([]int{1, 3, 4})))  // [1,1,2,3,4,4]
	fmt.Println(mergeTwoLists2(transfrom([]int{1, 2, 4}), transfrom([]int{1, 3, 4}))) // [1,1,2,3,4,4]
}
