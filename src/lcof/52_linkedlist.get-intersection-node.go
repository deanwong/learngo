package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	lookup := make(map[*ListNode]bool)
	for a != nil {
		lookup[a] = true
		a = a.Next
	}
	for b != nil {
		if lookup[b] {
			return b
		}
		b = b.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	if a == nil || b == nil {
		return nil
	}
	// a + c + b = b + c + a  c 为公共链表长度
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
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
	fmt.Println(getIntersectionNode(transfrom([]int{4, 1, 8, 4, 5}), transfrom([]int{5, 0, 1, 8, 4, 5}))) // [8,4,5]
	fmt.Println(getIntersectionNode(transfrom([]int{2, 6, 4}), transfrom([]int{1, 5})))                   // nil
}
