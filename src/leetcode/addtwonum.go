package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		if tail == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func transfrom(num int) *ListNode {
	var head *ListNode
	var tail *ListNode
	n := num
	if n == 0 {
		return &ListNode{Val: 0}
	}
	for n > 0 {
		if tail == nil {
			head = &ListNode{Val: n % 10}
			tail = head
		} else {
			tail.Next = &ListNode{Val: n % 10}
			tail = tail.Next
		}
		n = n / 10
	}
	return head
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
	fmt.Print(addTwoNumbers(transfrom(342), transfrom(465)))
	fmt.Print(addTwoNumbers(transfrom(0), transfrom(0)))
	fmt.Print(addTwoNumbers(transfrom(9999999), transfrom(9999)))
}
