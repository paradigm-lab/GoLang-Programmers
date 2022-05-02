package main

import (
	"fmt"
	"strconv"
)

// Declaring variable at the package level
var p float32 = 42

// Declaring multiple variables
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {

	// Variable format in Golang

	// We use this format when we want to declare a variable
	// and we don't need to use the variable yet or Initialize
	var j int
	j = 27

	var i float32 = 42

	k := 99

	// %v means the value and %T means Type
	fmt.Printf("%v %T \n", j, j)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T \n", k, k)
	fmt.Printf("%v %T \n", p, p)

	fmt.Printf("\n \n \n")
	fmt.Printf("Type convertion of the variables from int to float \n")

	// Variable type converstion (int to float)
	var typeI int = 42
	fmt.Printf("%v %T \n", typeI, typeI)

	var typeJ float32
	typeJ = float32(typeI)
	fmt.Printf("%v %T \n", typeJ, typeJ)

	fmt.Printf("\n \n")
	fmt.Printf("Type convertion of the variables from float to int \n")

	// Variable type converstion (float to int)
	// But to keep in mind that a floating point number can have a decimal
	// Now we have lost the information through the converstion
	// Explicity convertion
	var typeF float32 = 42.29
	fmt.Printf("%v %T \n", typeF, typeF)

	var typeH int
	typeH = int(typeF)
	fmt.Printf("%v %T \n", typeH, typeH)

	fmt.Printf("\n \n")
	fmt.Printf("Type convertion of the variables from float to int \n")

	// Converating an integer to a String
	// We can't converte Integer to a String
	// Because String is an aliase of the stream of bits (UTF-8)
	var typeInt int = 42
	fmt.Printf("%v %T \n", typeInt, typeInt)

	var typeStr string
	typeStr = strconv.Itoa(typeInt)
	fmt.Printf("%v %T \n", typeStr, typeStr)

}
