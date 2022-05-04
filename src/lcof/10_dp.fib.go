package main

import "fmt"

func fib(n int) int {
	const mod int = 1e9 + 7
	if n == 0 || n == 1 {
		return n
	}
	a := 0
	b := 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (a + b) % mod
		a = b
		b = ans
	}
	return ans
}

func main() {
	fmt.Println(fib(2)) // 1
	fmt.Println(fib(5)) // 5
}
