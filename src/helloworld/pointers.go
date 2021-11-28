package main

import "fmt"

func main() {
	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	*name = "David"
}
