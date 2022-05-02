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

	/* 
	 "Syntactic Sugar" in Programming
		-> A term to describe a feature in a language that let's you do samething more easily
		-> Makes the language "sweeter for human use
		-> But doesn't add any new functionality that it didn't already have

	 We can not use := with constants
	*/

	conferenceName := "Go Conference"

	// Constants are like variables, except that their values cannot be changed
	const conferenceTickets int = 50

	var remainingTickets uint = 50	

	// Printing the Type for the variables Using the %T for the Type
	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	/* 
	   Data Types:
		-> In any Programming language you have multiple data types
		-> Difference: which data types do they support exactly?
	
	   String -> For textual data, defined with double quotes, eg: "This is a String"
	
	   Integers -> Representing whole number, positive and negative, eg: 5, 120, -20
		    -> There are many more numeric data types!
			Different Integer Types and Different Floating Types
		
	   -> Each data type can do different things and behaves differently

	   Go is a statically typed language
		-> You need to tell Go Compiler, the data type when declaring the variable
			-> More Robust, Reduce the likelihood of errors
			-> Helps developers to catch type mismatches sooner (at compile time)

		-> Type Inference: BUT, Go can infer the type when you assign a value using the assignment operator	   	
	*/

	var userName string
	var userTickets int

	/*
		 A pointer is a variable that point to them memory address of another variable
		 A pointer is also known as A Special variable
	*/

	// Ask user for their name 
	fmt.Scan(&userName)

	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
}
