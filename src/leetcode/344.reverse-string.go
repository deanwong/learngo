package main

import "fmt"

// O(N)
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		temp := s[j]
		s[j] = s[i]
		s[i] = temp
		i++
		j--
	}
}

func main() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(a)
	fmt.Println(a)
	b := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(b)
	fmt.Println(b)
}
