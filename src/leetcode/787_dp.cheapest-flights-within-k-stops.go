package main

import (
	"fmt"
	"math"
)

// O(2^n * n) 2^n 表示递归次数，n表示复制答案的次数
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	// definition: dp[i][k] 表示 从 src 到 i 经过 k个机场 的最低 price
	// formulation : dp[flight[0]][k-1] (即 src 到 flight[0] 经过 k-1 个机场的最低 price) 如果有值，那么 dp[flight[1]][k] (即 src 到 flight[1] 经过 k个机场的最低 price) = dp[flight[0]][k-1] + flight[2] 中的较小值
	// initial: dp[src]][k+1] = 0 (1<=k<=4) 即 src 到 src不管中转 k 次都是0, 其他都是无穷大
	// answer: dp[dst][k+1]
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+2)
		for j := 0; j <= k+1; j++ {
			if i == src {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt32
			}
		}
	}
	// 开始填表
	for _k := 1; _k <= k+1; _k++ {
		for j := 0; j < len(flights); j++ {
			flight := flights[j]
			if dp[flight[0]][_k-1] != math.MaxInt32 {
				dp[flight[1]][_k] = min(dp[flight[1]][_k], dp[flight[0]][_k-1]+flight[2])
			}
			fmt.Println(dp)
		}
	}
	fmt.Println(dp)
	if dp[dst][k+1] == math.MaxInt32 {
		return -1
	}
	return dp[dst][k+1]
}

func findCheapestPrice_dfs(n int, flights [][]int, src int, dst int, k int) int {
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var dfs func(i int, k int) int
	// 表示从 i 到 dst 的走 k 步的最小价格
	dfs = func(i int, k int) int {
		fmt.Printf("%v->%v, k %v\n", i, dst, k)
		price := math.MaxInt32
		if k < 0 {
			fmt.Printf("%v->%v, k %v, price %v\n", i, dst, k, price)
			return price
		}
		if i == dst {
			fmt.Printf("%v->%v, k %v, price %v\n", i, dst, k, 0)
			return 0
		}
		for idx, flight := range flights {
			if flight[0] == i {
				// 遍历 i 的下一个节点
				price = min(price, dfs(flight[1], k-1)+flight[2])
				fmt.Printf("[%v] %v->%v, k %v, cur %v, price %v\n", idx, i, dst, k, flight[2], price)
			}
		}
		return price
	}
	ans := dfs(src, k+1)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	/*
		+------+---+-----+-----+-----+
		| *    | k | 0   | 1   | 2   |
		+------+---+-----+-----+-----+
		| city | * | *   | *   | *   |
		+------+---+-----+-----+-----+
		| 0    | * | 0   | 0   | 0   |
		+------+---+-----+-----+-----+
		| 1    | * | MAX | 100 | 100 |
		+------+---+-----+-----+-----+
		| 2    | * | MAX | 500 | 200 |
		+------+---+-----+-----+-----+
		dp[2][2] 考查的是 dp[1][1] 即从0到1经过1个机场的最低价 , 即 100 + {1, 2, 100}[2] = 200
	*/
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)) // 200
	/*
		+------+---+-----+-----+
		| *    | k | 0   | 1   |
		+------+---+-----+-----+
		| city | * | *   | *   |
		+------+---+-----+-----+
		| 0    | * | 0   | 0   |
		+------+---+-----+-----+
		| 1    | * | MAX | 100 |
		+------+---+-----+-----+
		| 2    | * | MAX | 500 |
		+------+---+-----+-----+
	*/
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)) // 500

	fmt.Println(findCheapestPrice_dfs(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)) // 200
	fmt.Println(findCheapestPrice_dfs(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)) // 500
}
