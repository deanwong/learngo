package main

import (
	"fmt"
)

// 带size的最小堆，堆底是最大的数，对顶就是第k个最大的数
type Heap struct {
	arr  []int
	size int
}

// 堆化
func (hp *Heap) Heapify(nums []int) {
	if len(nums) > hp.size {
		nums = nums[:hp.size]
	}
	hp.arr = make([]int, hp.size)
	copy(hp.arr, nums)
	for i := len(hp.arr) / 2; i >= 0; i-- {
		hp.Down(i)
	}
}

// 超过堆顶(min)才可以加入堆
func (hp *Heap) Push(num int) {
	if num > hp.arr[0] {
		hp.arr[0] = num
		hp.Down(0)
	}
}

// func (hp *Heap) Add(num int) {
// 	if len(hp.arr) < hp.size {
// 		hp.arr = append(hp.arr, num)
// 		for i := len(hp.arr) - 1; i > 0; {
// 			p := (i - 1) / 2
// 			if p >= 0 && hp.arr[p] > hp.arr[i] {
// 				hp.Swap(p, i)
// 				i = p
// 			} else {
// 				break
// 			}
// 		}
// 	} else if num > hp.arr[0] {
// 		hp.arr[0] = num
// 		hp.Down(0)
// 	}
// }

func (hp *Heap) Down(i int) {
	l, r, min := i*2+1, i*2+2, i
	n := len(hp.arr)
	if l < n && hp.arr[l] < hp.arr[min] {
		min = l
	}
	if r < n && hp.arr[r] < hp.arr[min] {
		min = r
	}
	if min != i {
		hp.Swap(i, min)
		hp.Down(min)
	}
}

func (hp *Heap) Swap(a, b int) {
	hp.arr[a], hp.arr[b] = hp.arr[b], hp.arr[a]
}

func findKthLargest(nums []int, k int) int {
	hp := &Heap{size: k}
	hp.Heapify(nums)
	for i := k; i < len(nums); i++ {
		hp.Push(nums[i])
	}
	return hp.arr[0]
}

// func buildMaxHeap(a []int, heapSize int) {
// 	for i := heapSize / 2; i >= 0; i-- {
// 		maxHeapify(a, i, heapSize)
// 	}
// }

// func maxHeapify(a []int, i, heapSize int) {
// 	l, r, max := i*2+1, i*2+2, i
// 	if l < heapSize && a[l] > a[max] {
// 		max = l
// 	}
// 	if r < heapSize && a[r] > a[max] {
// 		max = r
// 	}
// 	if max != i {
// 		a[i], a[max] = a[max], a[i]
// 		maxHeapify(a, max, heapSize)
// 	}
// }

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // 5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // 4
}
