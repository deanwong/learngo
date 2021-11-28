package main

import "fmt"

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n <= 0 {
		return ans
	}
	dfs(&ans, "", 0, 0, n)
	return ans
}

func dfs(ans *[]string, str string, left int, right int, n int) {
	if len(str) == n*2 {
		*ans = append(*ans, str)
		return
	}
	if left < n {
		str += "("
		dfs(ans, str, left+1, right, n)
		str = str[:len(str)-1]
	}
	if right < left {
		str += ")"
		dfs(ans, str, left, right+1, n)
		str = str[:len(str)-1]
	}
}

func main() {

	fmt.Println(generateParenthesis(0))
}
