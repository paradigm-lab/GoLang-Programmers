package main 

import (
	"fmt"
)

func main() {
	// An array has got a fixed value and one type
	grades := [...]int{97, 85, 93}
	var students [3]string
	
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Students: %v \n", students)

	// Assigning the values into the array using the index 
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v \n", students)

	// Getting the length of the array
	fmt.Printf("Number of Students: %v \n", len(students))

	// Arrays of arrays
	var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1} }
	
	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 1, 0}
	matrix[2] = [3]int{0, 0, 1}

	fmt.Printf("\n")

	fmt.Println(matrix)
	
	fmt.Printf("\n")

	fmt.Println(identityMatrix)
	
	fmt.Printf("\n")

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", a, a)
	
	x := [...]int{1, 2, 3}
	y := &x
	y[1] = 5
	
	fmt.Printf("\n")

	fmt.Printf("%v, %T \n", x, x)
	fmt.Printf("%v, %T \n", y, y)
	fmt.Printf("%v, %T \n", x, x)
}
