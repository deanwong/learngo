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

func findKthLargest(nums []int, k int) int {
	h := make(maxHeap, len(nums))
	copy(h, maxHeap(nums))
	heap.Init(&h)
	for i := k - 1; i > 0; i-- {
		heap.Pop(&h)
	}
	return h[0]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // 5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // 4
}
