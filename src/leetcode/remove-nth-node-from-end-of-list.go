package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	if head.Next == nil {
		return nil
	}
	for fast != nil {
		fast = fast.Next
		// last node
		if fast == nil {
			slow.Next = slow.Next.Next
		}
		n--
		if n <= 0 {
			slow = slow.Next
		}
	}
	return dummy.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
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
			tail = &ListNode{Val: n % 10}
			head = tail
		} else {
			head = &ListNode{Val: n % 10}
			head.Next = tail
			tail = head
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
	fmt.Println(removeNthFromEnd(transfrom(12345), 2))
	fmt.Println(removeNthFromEnd(transfrom(1), 1))
	fmt.Print(removeNthFromEnd(transfrom(12), 2))
}
