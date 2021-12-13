package main

import "fmt"

// O(2^n * n) 2^n 表示递归次数，n表示复制答案的次数
func letterCasePermutation(s string) []string {
	ans := make([]string, 0)
	dfs(&ans, s, 0, "")
	return ans
}

func dfs(ans *[]string, s string, idx int, combine string) {
	if idx == len(s) {
		*ans = append(*ans, combine)
		return
	}
	ch := s[idx]
	choices := make([]byte, 0)
	if ch >= 65 && ch <= 90 {
		choices = append(choices, ch+32)
	} else if ch >= 97 && ch <= 122 {
		choices = append(choices, ch-32)
	}
	choices = append(choices, ch)
	for i := 0; i < len(choices); i++ {
		combine = combine + string(choices[i])
		dfs(ans, s, idx+1, combine)
		combine = combine[:len(combine)-1]
	}
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("12345"))
	fmt.Println(letterCasePermutation("0"))
	fmt.Println(letterCasePermutation(""))
}
