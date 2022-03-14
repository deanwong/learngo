package main

import "fmt"

type MaxQueue struct {
	// 按时间序
	q []int
	// 单调递减
	p []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.p[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	// 大于队末数值，那么删除队尾元素
	for len(this.p) > 0 && value > this.p[len(this.p)-1] {
		this.p = this.p[:len(this.p)-1]
	}
	this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	// 队首元素如何恰好也是单调队列队首元数，需要调整单调队列
	if this.p[0] == this.q[0] {
		this.p = this.p[1:]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}

func main() {
	obj := Constructor()
	obj.Push_back(1)
	obj.Push_back(2)
	fmt.Println(obj.Max_value())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Max_value())
}
