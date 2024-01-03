package main

import (
	"fmt"
	"sync"
	"time"
)

// making the most used variables package level
// declariing variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

// instead of using a map we can use a struct which allows data of different types
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// creating a wait group to synchronize the go routine
var wg = sync.WaitGroup{}

func main() {
	// calling the greeting function
	greetUser()

	// %T is used to get the type of the data
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	// using an infinate for loop in Go

	// calling the getUserInput function
	firstName, lastName, email, userTickets := getUserInput()

	// calling the validate user inputs
	// we're using helper.Validate... because the function is from another package
	isValidUserName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// Ensuring that remaining tickets are more than userTickets
	if isValidUserName && isValidEmail && isValidTicket {

		// calling the booking ticket function
		bookTicket(firstName, lastName, email, userTickets)

		// sending ticket to the recepient. Initializing the go routine for concurrency
		// adding the add method from the wait group
		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		//Getting the firstNames of the users
		firstNames := []string{}
		firstNames = getFirstNames(firstNames)

		// Checking the remaining tickets
		if remainingTickets == 0 {
			fmt.Printf("All tickets for the %v are sold out.\n", conferenceName)
		}

		// calling the conference details
		getConfDetails(firstName, lastName, email, userTickets, firstNames)

	} else {
		if !isValidUserName {
			fmt.Println("Your first and last name are less than two characters.")
		}
		if !isValidEmail {
			fmt.Println("Your email address doesn't contain @gmail.com sign.")
		}
		if !isValidTicket {
			fmt.Println("You have booked many tickets than available ones.")
		}
	}

	// wait method to wait for all the threads that were add 
	wg.Wait()
}

// creating a greeting user function
func greetUser() {
	// %v is used to get the value of the data
	fmt.Printf("Welcome to %v our booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

// function to get user input
func getUserInput() (string, string, string, uint) {
	// Getting user input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// we're using &firstName to get the address of the memory so that we can be able it's location.
	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets purchased : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// geting the conference details
func getConfDetails(firstName string, lastName string, email string, userTickets uint, firstNames []string) {
	fmt.Printf("The first slice value : %v\n", bookings[0])
	fmt.Printf("The slice length : %v.\n", len(bookings))
	fmt.Printf("slice type : %T.\n", bookings)

	fmt.Printf("Dear %v %v you have successfully booked %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v.\n", remainingTickets, conferenceName)
	fmt.Printf("These are all the firstNames of bookings %v.\n", firstNames)
}

// geting all the first names
func getFirstNames(firstNames []string) []string {
	for _, booking := range bookings {
		// taking the firstName of the two names at index 0
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// booking ticket function
func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	// Updating the number of remaining tickets
	remainingTickets = remainingTickets - userTickets

	// creating a userData struct
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// Inputing users into the array
	bookings = append(bookings, userData)
	fmt.Printf("User details in a struct %v.\n", userData)
}

// function to send ticket to the email to simulate concurrency
func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	// simulating a large number of users signing in
	time.Sleep(50 * time.Second)
	// sending the tickets to the email
	var ticket = fmt.Sprintf("%v tickets sent to %v %v.\n", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending tickets:\n %v email address %v.\n", ticket, email)
	fmt.Println("######################")
	// adding the done method which removes the added go routines from the waiting list
	wg.Done()
}
