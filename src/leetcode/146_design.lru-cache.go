package main

import "fmt"

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity, cache: make(map[int]*DLinkedNode)}
	lru.head = &DLinkedNode{key: 0, value: 0}
	lru.tail = &DLinkedNode{key: 0, value: 0}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		// 移到最前
		this.moveToHead(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.value = value
		// 移到最前
		this.moveToHead(v)
	} else {
		node := &DLinkedNode{key: key, value: value}
		this.cache[key] = node
		// 直接加入到最前
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			delete(this.cache, this.removeTail().key)
			this.size--
		}
	}
}

func (this *LRUCache) deleteNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.deleteNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1)) // 1
	obj.Put(3, 3)
	fmt.Println(obj.Get(2)) // -1
	obj.Put(4, 4)
	fmt.Println(obj.Get(1)) // -1
	fmt.Println(obj.Get(3)) // 3
	fmt.Println(obj.Get(4)) // 4
}
