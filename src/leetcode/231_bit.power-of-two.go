package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 != 0 {
		return false
	}
	for n >= 2 {
		n = n / 2
		return isPowerOfTwo(n)
	}
	return false
}

func isPowerOfTwo_plus(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
}
