package main

// Formart package or fmt
import (
	"fmt" 
	"strings"
)

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

	
	// Data structure to store collection of elements in a single variable
	// Using Arrays and slices
	// Arrays in Go have got the Fixed size (size = how many elements the array can hold) and data type or (type)
	// var bookings = [50]string {}


	// A list that is more dynamic in size 
	// Slice is an abstraction of an Array
	// Slice are more flexible and powerfull:
	// Variable-length or get an sub-array of its own
	// Slices are also index-based and have a size, but is resized when needed
	 var bookings []string


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

	for {
		
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		/*
			 A pointer is a variable that point to them memory address of another variable
			 A pointer is also known as A Special variable
		*/

		// Ask user for their name 
		fmt.Println("Enter your firstName: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastName: ")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter your userTickets: ")
		fmt.Scan(&userTickets)


		// Validation
		// Logical Operators: &&, ||, !=
		// First Name and Last Name validation
		isValidName := len(firstName) >= 2 &&  len(lastName) >= 2

		// Email validation
		isValidEmail := strings.Contains(email, "@")
		
		// Ticket Validation (Needs to enter correct number of ticket positive + greater than 0)
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets 


		// City input using the switch statementa using the multiple value to execute same logic



		// Logic to test for the user input if it's greater than then remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {   
			// Calculation in go should be done with the same type
			// To handle this issue is to convert one of them to the other type using the builtIn function that are available in the language
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName + " " + lastName)	

			/*
			fmt.Printf("The whole array: %v \n", bookings)
			fmt.Printf("The first value: %v \n", bookings[0])
			fmt.Printf("Array type: %T \n", bookings)
			fmt.Printf("Array length: %v \n", len(bookings))
			*/
			
			/*
			fmt.Printf("The whole array: %v \n", bookings)
			fmt.Printf("The first value: %v \n", bookings[0])
			fmt.Printf("Slice type: %T \n", bookings)
			fmt.Printf("Slice length: %v \n", len(bookings))
			*/
			
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email) 	
			fmt.Printf("%v ticket remaining for %v \n", remainingTickets, conferenceName)	
			
			firstNames := []string{}

			// Range iterates over elements for different data structure (so not only arrys and slices)
			// Blank Identifier _ 
			for _, booking := range bookings {
				// strings.Fields() Splits the strings with white space as separator and returns a slice with the split element
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("The first names of bookings are: %v \n", firstNames)
			
			var noTicketsRemaining bool = remainingTickets == 0 

			if noTicketsRemaining { 
				// end program
				fmt.Println("OUr conference is booked out. Come back next yeare.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}


			// fmt.Printf("Your input data is invalid, try again \n")
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
			// continue statement causes loop to skep the remainder of its body and immediately retesting its condition (in our infinite loop our condition is always true)
			// continue
		}
	}

}
