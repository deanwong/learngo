package main

import (
	"fmt"
	"math/rand"
)

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

type Solution struct {
	origin []int
	nums   []int
}

func Constructor(nums []int) Solution {
	clone := make([]int, len(nums))
	copy(clone, nums)
	s := Solution{origin: clone, nums: nums}
	return s
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.origin)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	for i := n - 1; i > 0; i-- {
		// 从后往 [0: i] 替换，注意必须包含i
		// i=n-1时，有n种选择[0:n-1]，i=n-2时，有n-1中选择[0:n-2]
		j := rand.Intn(i + 1)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	// 也可以从前往后 i=0时，有n种选择[0:n-1]，i=1时，有n-1中选择[1:n-1]
	// for i := 0; i < n; i++ {
	// 	swap(i, i+rand.Intn(n-i))
	// }
	return this.nums
}

func (this Solution) String() string {
	return fmt.Sprintf("origin %v, nums %v", this.origin, this.nums)
}

func main() {
	obj := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	obj.Shuffle()
	fmt.Println(obj)
	obj.Reset()
	fmt.Println(obj)
	obj.Shuffle()
	fmt.Println(obj)
	obj.Reset()
	fmt.Println(obj)
}
