package main 


import (
	"fmt"
)


func main() {
	const myConst int = 42 // Internalle constant
	// Export constant const MyConst
	fmt.Printf("%v, %T\n", myConst, myConst)

}
