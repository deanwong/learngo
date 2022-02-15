package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	// definition: dp[i] 表示前 i 个数字（实际是i-1）有几种翻译方法
	// formulation :
	// 1. s[i]>=1 && s[i]<=9 只能有位置 i 组成一种方式 dp[i] = dp[i-1]
	// 2. s[i-1] *10 + s[i]>=10 && s[i-1] *10 + s[i]<=25 只能当前位置 i 的与前一位置（i-1）组成 dp[i] = dp[i-2]
	// initial: dp[0] 只有一种翻译方法
	// answer: dp[n-1]
	str := strconv.Itoa(num)
	n := len(str)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		digit := int(str[i] - '0')
		twoDigits := int(str[i-1]-'0')*10 + digit
		if twoDigits >= 10 && twoDigits <= 25 {
			if i >= 2 {
				dp[i] += dp[i-2]
			} else {
				dp[i]++
			}
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(translateNum(12258)) // 5
}
