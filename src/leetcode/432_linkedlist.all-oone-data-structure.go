package main

import (
	"container/list"
	"fmt"
)

type linkedNode struct {
	keys  map[string]int // 出现count次数的key的集合
	count int            // 出现了count次
}

type AllOne struct {
	*list.List                          //  双向链表
	nodes      map[string]*list.Element // 哈希表
}

func Constructor() AllOne {
	return AllOne{list.New(), make(map[string]*list.Element)}
}

func (this *AllOne) Inc(key string) {
	cur := this.nodes[key]
	// 当前节点已存在
	if cur != nil {
		curNode := cur.Value.(linkedNode)
		next := cur.Next()
		// 判断当前节点与下个节点是否有插入的空隙
		if next == nil || next.Value.(linkedNode).count > curNode.count+1 {
			this.nodes[key] = this.InsertAfter(linkedNode{map[string]int{key: 1}, curNode.count + 1}, cur)
		} else {
			next.Value.(linkedNode).keys[key] = 1
			this.nodes[key] = next
		}
		// 从当前节点中删除
		delete(curNode.keys, key)
		// 直接删除当前节点
		if len(curNode.keys) == 0 {
			this.Remove(cur)
		}
	} else {
		// 如果链表头节点为空(说明没有节点)或者是链表的头节点的count就大于1(因为新key进入插入的节点的count应该是1)，则在头节点前插入一个新node
		if this.Front() == nil || this.Front().Value.(linkedNode).count > 1 {
			this.nodes[key] = this.PushFront(linkedNode{map[string]int{key: 1}, 1})
		} else {
			// 在头节点的keys中加入key
			this.Front().Value.(linkedNode).keys[key] = 1
			this.nodes[key] = this.Front()
		}
	}
}

func (this *AllOne) Dec(key string) {
	cur := this.nodes[key]
	curNode := cur.Value.(linkedNode)
	if curNode.count > 1 {
		// 表示存在
		pre := cur.Prev()
		// 如果前面的节点和当前节点有空隙
		if pre == nil || pre.Value.(linkedNode).count < curNode.count-1 {
			this.nodes[key] = this.InsertBefore(linkedNode{map[string]int{key: 1}, curNode.count - 1}, cur)
		} else {
			// 在前置节点的node里存入key
			pre.Value.(linkedNode).keys[key] = 1
			this.nodes[key] = pre
		}
	} else {
		// 因为数量是1，在nodes中直接删除
		delete(this.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		this.Remove(cur)
	}
}

func (this *AllOne) GetMaxKey() string {
	// 最大值就是取尾节点
	if b := this.Back(); b != nil {
		for k := range b.Value.(linkedNode).keys {
			return k
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	// 最小值就是取头节点
	if f := this.Front(); f != nil {
		for k := range f.Value.(linkedNode).keys {
			return k
		}
	}
	return ""
}

func main() {
	allone := Constructor()
	allone.Inc("hello")
	allone.Inc("hello")
	fmt.Println(allone.GetMaxKey())
	fmt.Println(allone.GetMinKey())
	allone.Inc("leet")
	fmt.Println(allone.GetMaxKey())
	fmt.Println(allone.GetMinKey())
}
