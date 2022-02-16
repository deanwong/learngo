package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)

	rune := 'F'
	fmt.Println(rune)

	// var integer32 int = 2147483649
	// fmt.Println(integer32)

	// var integer uint = -10
	// fmt.Println(integer)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	var featureFlag bool = true
	fmt.Println(featureFlag)

	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)

	var integer16 int16 = 127
	var integer32 int32 = 32767
	fmt.Println(int32(integer16) + integer32)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
