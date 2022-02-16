package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
}

func sum(a string, b string) int {
	number1, _ := strconv.Atoi(a)
	number2, _ := strconv.Atoi(b)
	return number1 + number2
}
