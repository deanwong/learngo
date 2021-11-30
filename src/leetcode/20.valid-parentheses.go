package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{')': '(', '}': '{', ']': '['}

	var stack []byte
	for i := 0; i < n; i++ {
		c := s[i]
		if left, ok := pairs[c]; ok {
			if len(stack) < 1 || stack[len(stack)-1] != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// left part
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("})"))
}
