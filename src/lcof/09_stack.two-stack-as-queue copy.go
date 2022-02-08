package main

import "fmt"

type CQueue struct {
	stack        []int
	stackReverse []int
}

func Constructor() CQueue {
	return CQueue{make([]int, 0), make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.stack = append(this.stack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack) == 0 && len(this.stackReverse) == 0 {
		return -1
	} else if len(this.stackReverse) > 0 {
		// 什么
		val := this.stackReverse[len(this.stackReverse)-1]
		this.stackReverse = this.stackReverse[:len(this.stackReverse)-1]
		return val
	}
	// 对于golang来说没必要，模拟其他stack的特性
	for len(this.stack) > 0 {
		this.stackReverse = append(this.stackReverse, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
	val := this.stackReverse[len(this.stackReverse)-1]
	this.stackReverse = this.stackReverse[:len(this.stackReverse)-1]
	return val
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	obj1 := Constructor()
	obj1.AppendTail(3)
	fmt.Println(obj1.DeleteHead())
	fmt.Println(obj1.DeleteHead())

	obj2 := Constructor()
	fmt.Println(obj2.DeleteHead())
	obj2.AppendTail(5)
	obj2.AppendTail(2)
	fmt.Println(obj2.DeleteHead())
	fmt.Println(obj2.DeleteHead())
}
