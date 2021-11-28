package main

import "fmt"

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			panic("negative number")
		case val == 0:
			fmt.Prinln("0 is neither negative nor positive")
		case val > 0:
			fmt.Println("You entered:", val)
		}
	}
}
