package main

import (
	"fmt"
)

func isPrimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	if number == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		if isPrimes(number) {
			fmt.Printf("Prime %v \n", number)
		}
	}
}
