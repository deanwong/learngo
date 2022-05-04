package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 旧节点 -> 新节点
	dict := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		dict[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		dict[cur].Next = dict[cur.Next]
		dict[cur].Random = dict[cur.Random]
		cur = cur.Next
	}
	return dict[head]
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. 复制各节点，并构建拼接链表
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}
	// 2. 构建各新节点的 random 指向
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 3. 拆分两链表
	pre := head     // 原链表
	cur = head.Next // 新链表
	ans := head.Next
	for cur.Next != nil {
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil
	return ans
}

func transfrom(val [][2]int) *Node {
	if val == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	dummy := &Node{Val: 0}
	cur := dummy
	for i := 0; i < len(val); i++ {
		node := &Node{Val: val[i][0]}
		nodes = append(nodes, node)
		cur.Next = node
		cur = cur.Next
	}
	cur = dummy.Next
	for i := 0; i < len(val); i++ {
		if val[i][1] >= 0 {
			cur.Random = nodes[val[i][1]]
		}
		cur = cur.Next
	}
	return dummy.Next
}

func (head *Node) String() string {
	val := make([][2]int, 0)
	cur := head
	for cur != nil {
		random := -1
		if cur.Random != nil {
			random = cur.Random.Val
		}
		val = append(val, [2]int{cur.Val, random})
		cur = cur.Next
	}
	return fmt.Sprintf("%v", val)
}

func main() {
	fmt.Println(transfrom([][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}))
}
