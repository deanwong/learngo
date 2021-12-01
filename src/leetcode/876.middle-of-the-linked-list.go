package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n)
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	step := 0
	for fast != nil {
		fast = fast.Next
		step++
		if step%2 == 0 {
			slow = slow.Next
		}
	}
	return slow
}

func transfrom(num int) *ListNode {
	dummy := &ListNode{Val: 0}
	n := num
	if n == 0 {
		return &ListNode{Val: 0}
	}
	for n > 0 {
		node := &ListNode{Val: n % 10, Next: dummy.Next}
		dummy.Next = node
		n = n / 10
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
	fmt.Print(middleNode(transfrom(123456)))
}
