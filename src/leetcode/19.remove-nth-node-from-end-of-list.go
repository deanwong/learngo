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
// O(L)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	if head.Next == nil {
		return nil
	}
	for fast != nil {
		fast = fast.Next
		n--
		if n < 0 {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
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
	fmt.Println(removeNthFromEnd(transfrom(12345), 2))
	fmt.Println(removeNthFromEnd(transfrom(1), 1))
	fmt.Print(removeNthFromEnd(transfrom(12), 2))
}
