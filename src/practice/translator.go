package main

import (
	"fmt"
)

func main() {
	sum := translator("MCMZ")
	fmt.Println(sum)
}

func translator(word string) int {
	romandict := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	numbers := make([]int, len(word))
	for i, c := range word {
		if val, present := romandict[c]; present {
			numbers[i] = val
		} else {
			return 0
		}
	}
	fmt.Println(numbers)

	sum := 0

	for i := 0; i < len(word); i++ {
		if i != len(word)-1 && numbers[i] < numbers[i+1] {
			numbers[i] = -numbers[i]
		}
		sum += numbers[i]
	}
	return sum
}
