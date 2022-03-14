package main

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
	for fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow.val
}

func main() {
	root := &LinkedNode{val: 1}
	root.next = &LinkedNode{val: 2}
	root.next.next = &LinkedNode{val: 3}

}
