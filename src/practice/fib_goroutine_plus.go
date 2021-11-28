package main

import (
	"fmt"
	"time"
)

func fib(ch chan float64) {
	x, y := 1.0, 1.0
	for {
		x, y = y, x+y
		ch <- x
	}

}

func main() {
	start := time.Now()

	command := ""

	ch := make(chan float64)
	go fib(ch)

	for {
		fmt.Scanf("%s", &command)
		fmt.Printf("%v\n", <-ch)
		if command == "quit" {
			fmt.Println("Done calculating Fibonacci!")
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
