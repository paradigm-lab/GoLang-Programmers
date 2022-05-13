package main 


import (
	"fmt"
)


// iota is scoped to a constant block
const (
	p = iota 
	o = iota
	j = iota
)

const (
	p2 = iota 
)

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_ = iota // Ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	const myConst int = 42 // Internalle constant
	// Export constant const MyConst

	// We are not allowed to change the value of the constant
	// It has to be assinable/determined at compile time

	fmt.Printf("%v, %T\n", myConst, myConst)

	const a = 42
	var b int16 = 27

	// We can do implicity converstion with the constants but not with the nomarl variables
	fmt.Printf("%v, %T \n", a + b, a + b)

	fmt.Printf("%v, %T \n", p, p)
	fmt.Printf("%v, %T \n", o, o)
	fmt.Printf("%v, %T \n", j, j)

	fmt.Printf("\n")

	fmt.Printf("%v, %T \n", p2, p2)
	
	fmt.Printf("\n")

	var specialistType int
	fmt.Printf("%v \n", specialistType == catSpecialist)
	fmt.Printf("%v \n", catSpecialist)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}




