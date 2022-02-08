package main

import "fmt"

type MinStack struct {
	stack  []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.helper) == 0 {
		this.helper = append(this.helper, x)
		return
	}
	min := this.helper[len(this.helper)-1]
	if min > x {
		this.helper = append(this.helper, x)
	} else {
		this.helper = append(this.helper, min)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.helper = this.helper[:len(this.helper)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.helper) == 0 {
		return -1
	}
	return this.helper[len(this.helper)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
func main() {
	obj1 := Constructor()
	obj1.Push(-2)
	obj1.Push(0)
	obj1.Push(-3)
	fmt.Println(obj1.Min())
	obj1.Pop()
	fmt.Println(obj1.Top())
	fmt.Println(obj1.Min())
}
