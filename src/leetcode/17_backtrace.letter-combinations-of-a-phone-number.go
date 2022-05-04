package main

import "fmt"

var pairs = map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

// O(3^m×4^n)，其中 m 是输入中对应 3 个字母的数字个数（包括数字 m 是输入中对应 4 个字母的数字个数（包括数字 7、9），m+n 是输入数字的总个数
func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	if len(digits) <= 0 {
		return ans
	}
	var backtrace func(idx int, temp string)
	backtrace = func(idx int, temp string) {
		if idx == len(digits) {
			ans = append(ans, temp)
			return
		}
		ch := digits[idx]
		if letters, ok := pairs[ch]; ok {
			for i := 0; i < len(letters); i++ {
				temp = temp + string(letters[i])
				backtrace(idx+1, temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
	backtrace(0, "")
	return ans
}

func main() {
	fmt.Println(letterCombinations("23")) // [ad ae af bd be bf cd ce cf]
	fmt.Println(letterCombinations(""))   // []
	fmt.Println(letterCombinations("2"))  // [a b c]
}
