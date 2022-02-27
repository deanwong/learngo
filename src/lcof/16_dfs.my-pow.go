package main

import "fmt"

/*
假设我们要求一个数的32次方，如果我们已经知道了它的16次方，那么只要在16次方的基础上再平方一次就ok了；而16次方是8次方的平方 ；依此类推，我们求32次方只需要做5次乘法，先求平方，在平方的基础上求4次方，在4次方的基础上求8次方，在8次方的基础上求16次方，最后在16次方的基础上求32次方
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	result := myPow(x, n/2)
	result *= result
	// 如果是奇数次方，需要补1次
	if n%2 == 1 {
		result *= x
	}
	return result
}

func main() {
	// fmt.Println(myPow(2.00000, 10))  // 1024.0
	// fmt.Println(myPow(2.10000, 3))   // 9.26100
	// fmt.Println(myPow(2.00000, -2))  // 0.25000
	fmt.Println(myPow(34.00515, -3)) // 0.25000
}
