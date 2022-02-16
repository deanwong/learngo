package main

import (
	"fmt"

	"github.com/deanwong/learngo/calculator"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!")
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
	fmt.Println(quote.Hello())
}
