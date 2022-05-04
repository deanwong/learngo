package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *maxHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *maxHeap) Peek() int          { return (*h)[0] }

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *minHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *minHeap) Peek() int          { return (*h)[0] }

type MedianFinder struct {
	maxH *maxHeap
	minH *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{&maxHeap{}, &minHeap{}}
}

// 压入数据时有两种情况；
// 原有数据为奇数个，将数据放入 minHeap, 再从排序后的 minHeap pop 出最小值放入 maxHeap, 插入后两堆长度相等
// 上面做法的高明之处 在于 省去了判断 num 与 minHeap 最小值 及 maxHeap 最大值 的比较
// 原有数据为偶数个，将数据放入 maxHeap, 再从排序后的 maxHeap pop 出最小值放入 minHeap, 插入后 minHeap 比 maxHeap 长度大1
func (this *MedianFinder) AddNum(num int) {
	if this.maxH.Len() != this.minH.Len() {
		heap.Push(this.minH, num)
		heap.Push(this.maxH, heap.Pop(this.minH))
	} else {
		heap.Push(this.maxH, num)
		heap.Push(this.minH, heap.Pop(this.maxH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxH.Len() != this.minH.Len() {
		return float64(this.minH.Peek())
	} else {
		return float64(this.minH.Peek()+this.maxH.Peek()) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian()) // 1.5
	obj.AddNum(3)
	fmt.Println(obj.FindMedian()) // 2
}
