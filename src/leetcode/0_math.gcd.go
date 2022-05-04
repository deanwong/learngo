package main

import "fmt"

// O(n)
func gcd(a int, b int) int {
	fmt.Printf("a %v b %v\n", a, b)
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(gcd(18, 23))
	fmt.Println(gcd(54, 27))
}
