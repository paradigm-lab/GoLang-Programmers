package main

// Formart package or fmt
import "fmt" 

// Where to start the program? Where is the entrypoint
// The "main" function is the entrypoint of a Go Program

func main() {
	/*
	 Variables are used to store values
	 Like containers for values 
	 Give the variable a name and reference everywhere in the app
	 Variable names should be descriptive
	*/
	

	var conferenceName = "Go Conference"

	// Constants are like variables, except that their values cannot be changed
	const conferenceTickets = 50
	var remainingTickets = 50	

	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")


}
