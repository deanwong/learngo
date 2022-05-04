package main

import "fmt"

// 每个答案需要 O(n) 的时间复制到答案数组中
// O(n)
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n <= 0 {
		return ans
	}
	var backtrace func(left int, right int, temp string)
	backtrace = func(left int, right int, temp string) {
		if len(temp) == 2*n {
			ans = append(ans, temp)
			return
		}
		if left < n {
			temp = temp + "("
			backtrace(left+1, right, temp)
			temp = temp[:len(temp)-1]
		}
		if right < left {
			temp = temp + ")"
			backtrace(left, right+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backtrace(0, 0, "")
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3)) // [((())) (()()) (())() ()(()) ()()()]
	fmt.Println(generateParenthesis(1)) // [()]
}
