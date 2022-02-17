package main

import "fmt"

func reverseWords(s string) string {
	n := len(s)
	ans := ""
	l, r := 0, n-1
	for l < n && s[l] == ' ' {
		l++
	}
	for r >= 0 && s[r] == ' ' {
		r--
	}
	j := r + 1
	for i := r; i >= l; i-- {
		if s[i] == ' ' {
			ans = ans + s[i+1:j] + " "
			for i >= 1 && s[i-1] == ' ' {
				i--
			}
			j = i
		} else if i == l {
			ans += s[i:j]
		}
	}
	return ans
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))  // "blue is sky the"
	fmt.Println(reverseWords("  hello world!  ")) // "world! hello"
	fmt.Println(reverseWords("a good   example")) // "example good a"
}
