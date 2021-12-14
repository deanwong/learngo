package main

import "fmt"

func letterCombinations(digits string) []string {
	pairs := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	ans := make([]string, 0)
	if len(digits) <= 0 {
		return ans
	}
	dfs(pairs, &ans, digits, 0, "")
	return ans
}

func dfs(pairs map[byte]string, ans *[]string, digits string, index int, str string) {
	if len(digits) == index {
		*ans = append(*ans, str)
		return
	}
	digit := digits[index]
	for _, c := range pairs[digit] {
		str = str + string(c)
		dfs(pairs, ans, digits, index+1, str)
		str = str[:len(str)-1]
	}
}

func main() {
	fmt.Println(letterCombinations(""))
}
