package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	ans := make([]int, 0)
	for pre != nil {
		ans = append(ans, pre.Val)
		pre = pre.Next
	}
	return ans
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
	fmt.Println(reversePrint(transfrom(12345)))
}
