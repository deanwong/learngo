package main

import "fmt"

type LinkedNode struct {
	val  int
	next *LinkedNode
}

func (l *LinkedNode) push(node *LinkedNode) {
	l.next = node
}

func middle(root *LinkedNode) int {
	fast := root
	slow := root
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow.val
}

func main() {
	root := &LinkedNode{val: 1}
	root.push(&LinkedNode{val: 2})
	root.next.push(&LinkedNode{val: 3})
	fmt.Println(middle(root))
}
