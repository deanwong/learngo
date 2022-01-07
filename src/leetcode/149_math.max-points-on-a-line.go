package main

import (
	"fmt"
)

// O(n^3)
func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	// 枚举三点两条线的斜率
	// 两条直接的斜率满足 x1/x2==y1/y2 即 x1*y2==x2*y1 则重合
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			count := 2
			x1 := points[j][0] - points[i][0]
			y1 := points[j][1] - points[i][1]
			for k := j + 1; k < n; k++ {
				x2 := points[k][0] - points[j][0]
				y2 := points[k][1] - points[j][1]
				if x1*y2 == x2*y1 {
					count++
				}
			}
			ans = max(ans, count)
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))                         // 3
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}})) // 4
	fmt.Println(maxPoints([][]int{{0, 0}}))                                         // 1
}
