package main

import "fmt"

// O(s+t)
func backspaceCompare(s string, t string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	s = string(stack)
	stack = make([]byte, 0)
	for i := 0; i < len(t); i++ {
		if t[i] != '#' {
			stack = append(stack, t[i])
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	t = string(stack)
	return s == t
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("ab##", "c#d#")) // true
	fmt.Println(backspaceCompare("a##c", "#a#c")) // true
	fmt.Println(backspaceCompare("a#c", "b"))     // false
}
