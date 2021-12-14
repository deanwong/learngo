package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n)
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, dummy.Next
	for fast != nil && fast.Next != nil {
		if slow.Next.Val != fast.Next.Val {
			slow = slow.Next
		} else {
			for fast.Next != nil && fast.Next.Val == slow.Next.Val {
				fast = fast.Next
			}
			slow.Next = fast.Next
		}
		fmt.Printf("slow %v fast %v\n", slow.Val, fast.Val)
		fast = fast.Next
	}
	return dummy.Next
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
	fmt.Println(deleteDuplicates(transfrom(12334455))) // [1 2]
	fmt.Println(deleteDuplicates(transfrom(11123)))    // [2 3]
}
