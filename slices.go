package main 

import (
	"fmt"
)

func main() {

	// Slices are naturally called a reference type, Because the reffere to the same underlying data
	// Slices also use the index

	a := []int{1, 2, 3}
	
	b := a

	b[1] = 5

	fmt.Printf("%v %T \n", a, a)
	
	fmt.Printf("%v %T \n", b, b)

	// Length of slices
	fmt.Printf("Length: %v \n", len(a))

	// Capacity
	fmt.Printf("Capacity: %v \n", cap(a))

	fmt.Printf("\n")

	u := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v := u[:]  // Slice of all elements
	w := u[3:]  // Slice from 4th element to end
	x := u[:6]  // Slice first 6 elements
	y := u[3:6]  // Slice the 4th, 5th and 6th elements (inclusive:exclusive)

	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("\n")

	// Type, length and capacity
	f := make([]int, 3, 100)

	fmt.Println(f)

	// Length of slices
	fmt.Printf("Length: %v \n", len(f))

	// Capacity
	fmt.Printf("Capacity: %v \n", cap(f))

	fmt.Printf("\n")

	sliceAppend()

	stackpushpop()
}


func sliceAppend() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))

	// Create a new array 
	a = append(a, 1)
	
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
	
	a = append(a, 2, 3, 4, 5)
	
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
	
	fmt.Printf("\n")
}


// Concatenate a slice


func stackpushpop() {
	a := []int{1, 2, 3, 4, 5}
	// Removing or pop the last element in stack
	b := a[:len(a)-1]
	// Removing an element in the middle
	c := append(a[:2], a[3:]...) 

	fmt.Println(b)
	fmt.Println(c)
}
