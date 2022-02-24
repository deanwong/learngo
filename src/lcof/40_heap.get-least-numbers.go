package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

// 实现堆接口 的 五个方法
func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *maxHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	h := make(maxHeap, k)
	copy(h, maxHeap(arr[:k+1]))
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		if arr[i] < h[0] {
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}
	return h
}

func main() {
	// fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))    // [1,2]
	// fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1)) // [0]
	fmt.Println(getLeastNumbers([]int{0, 0, 0, 2, 0, 5}, 0))
}
