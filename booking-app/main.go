package main

// Formart package or fmt
import (
	"fmt" 
	"strings"
)


// Package level variable they are accessible to the main package
// Constants are like variables, except that their values cannot be changed
const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50	
var bookings = []string{}


// Where to start the program? Where is the entrypoint
// The "main" function is the entrypoint of a Go Program
func main() {


	// Function calling 
	greetUsers()
	
	// Arrays	
	// var bookings = [50]string {}

	for {
	
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)


		// Logic to test for the user input if it's greater than then remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {   
		
			//Function Call for the bookTicket 
			bookTicket(userTickets, firstName, lastName, email)

			// Call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)


			var noTicketsRemaining bool = remainingTickets == 0 

			if noTicketsRemaining { 
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}


		}  else {
			
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



func greetUsers() { 
	fmt.Printf("Welcome to %v booking application \n", conferenceName);
	// fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}


func getFirstNames() []string {
	
	// Range iterates over elements for different data structure (so not only arrys and slices)
	// Blank Identifier _ 
	firstNames := []string{}
	for _, booking := range bookings {
		// strings.Fields() Splits the strings with white space as separator and returns a slice with the split element
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}


func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// Validation
	// Logical Operators: &&, ||, !=
	// First Name and Last Name validation
	isValidName := len(firstName) >= 2 &&  len(lastName) >= 2

	// Email validation
	isValidEmail := strings.Contains(email, "@")

	// Ticket Validation (Needs to enter correct number of ticket positive + greater than 0)
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets 
	
	return isValidName, isValidEmail, isValidTicketNumber

	// City input using the switch statementa using the multiple value to execute same logic
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
} 

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Calculation in go should be done with the same type
	// To handle this issue is to convert one of them to the other type using the builtIn function that are available in the language
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)	


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email) 	
	fmt.Printf("%v ticket remaining for %v \n", remainingTickets, conferenceName)	
}
