package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	// 背包总价值为 w = sum / 2, 每一个物品只能装一次,如果出现背包中重量等于sum/2则为true 否则 false
	// definition: dp[i][j] 表示从数组的 [0, i] 这个子区间内挑选一些正整数，每个数只能用一次，使得这些数的和恰好等于 j
	// formulation: dp[i][j] = dp[i - 1][j] or dp[i - 1][j - nums[i]] 对于当前的第i个物品，有拿和不拿两种情况，dp[i - 1][j]表示不拿，dp[i - 1][j - nums[i]]表示拿
	// initial: dp[0][0]=0
	// answer: dp[n-1][w]
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 != 0 {
		return false
	}
	w := sum / 2
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, w+1)
	}
	// 物品栏第一行
	if nums[0] <= w {
		// 只能填充等于自己容量的背包
		dp[0][nums[0]] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			// 直接从上一行先把结果抄下来，然后再修正
			dp[i][j] = dp[i-1][j]
			// 单独 nums[j] 这个数恰好等于此时「背包的容积」 j
			if nums[i] == j {
				dp[i][j] = 1
				continue
			}
			// j - nums[i] 作为数组的下标，一定得保证大于等于 0 ，因此 nums[i] <= j
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1][w] != 0
}

func canPartition_plus(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 != 0 {
		return false
	}
	w := sum / 2
	// dp[j] 表示是否可以从数组中凑齐 j，即从物品栏凑出背包容量
	dp := make([]int, w+1)
	// 什么都不拿
	dp[0] = 1
	fmt.Println(dp)
	// 物品栏
	for i := 0; i < n; i++ {
		// 背包,  从后往前，j < nums[i] 马上退出当前循环，因为后面的 j 的值肯定越来越小
		for j := w; j >= nums[i]; j-- {
			// if dp[w] != 0 {
			// 	return true
			// }
			// dp[j] 表示不拿当前的物品，dp[j-nums[i]] 表示拿当前的物品所以要考查背包扣出当前物品价值时是否能填满
			dp[j] = dp[j] + dp[j-nums[i]]
		}
		fmt.Println(dp)
	}
	return dp[w] != 0
}

func main() {
	/*
	 *          0  1  2  3  4  5  6  7  8  9  10  11
	 *  nums[0] 0  1  0  0  0  0  0  0  0  0   0   0
	 *  nums[1] 0  1  0  0  0  1  1  0  0  0   0   0
	 *  nums[2] 0  1  0  0  0  1  1  0  0  0   0   1
	 *  nums[3] 0  1  0  0  0  1  1  0  0  0   0   1
	 */
	fmt.Println(canPartition([]int{1, 5, 11, 5}))      // true
	fmt.Println(canPartition_plus([]int{11, 5, 1, 5})) // true
	fmt.Println(canPartition([]int{1, 2, 3, 5}))       // false
}
