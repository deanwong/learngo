package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	deque := make([]int, 0)
	i := 1 - k
	for j := 0; j < len(nums); j++ {
		// 若 i > 0 且 队首元素最大元素就是被删除的元素，那么要从队首删除
		if i > 0 && deque[0] == nums[i-1] {
			deque = deque[1:]
		}
		// 队列中元素小于滑动窗口加入的数字则从队尾去掉
		for len(deque) > 0 && deque[len(deque)-1] < nums[j] {
			deque = deque[:len(deque)-1]
		}
		// 加入队列
		deque = append(deque, nums[j])
		// 形成窗口后，将队列首元素加入答案
		if i >= 0 {
			ans = append(ans, deque[0])
		}
		// 缩小左边界
		i++
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3,3,5,5,6,7]
}
