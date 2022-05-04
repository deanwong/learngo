package main

import (
	"fmt"
)

// O(2^n * n) 2^n 表示递归次数，n表示复制答案的次数
func letterCasePermutation(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	var backtrace func(start int, permutation string)
	backtrace = func(start int, permutation string) {
		if start == n {
			ans = append(ans, permutation)
			return
		}
		ch := s[start]
		choices := make([]byte, 0)
		choices = append(choices, ch)
		if ch >= 65 && ch <= 90 {
			choices = append(choices, ch+32)
		} else if ch >= 97 && ch <= 122 {
			choices = append(choices, ch-32)
		}
		for i := 0; i < len(choices); i++ {
			permutation = permutation + string(choices[i])
			backtrace(start+1, permutation)
			permutation = permutation[:len(permutation)-1]
		}

	}
	backtrace(0, "")
	return ans
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))  // [a1b2 a1B2 A1b2 A1B2]
	fmt.Println(letterCasePermutation("12345")) // [12345]
	fmt.Println(letterCasePermutation("0"))     // [0]
	fmt.Println(letterCasePermutation(""))      // []
}
