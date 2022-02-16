package main

import (
	"fmt"

	"github.com/deanwong/learngo/store"
)

func main() {
	tom, _ := store.NewEmployee("Tom", "Guy", 50.5)
	fmt.Println(tom.CheckCredits())
	fmt.Println(tom.AddCredits(20.1))
	c, _ := tom.RemoveCredits(10.6)
	fmt.Println(c)
	c, err := tom.RemoveCredits(100.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	tom.ChangeName("Lala")
	fmt.Println(tom)
}
