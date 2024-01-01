package main

import (
	"fmt"
	"strings"
)

func main() {
	// declariing variables
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	// %v is used to get the value of the data
	fmt.Printf("Welcome to %v our booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	// %T is used to get the type of the data
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	// using an infinate for loop in Go
	for {
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

		// validating user input
		isValidUserName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && remainingTickets >= userTickets

		// Ensuring that remaining tickets are more than userTickets
		if isValidUserName && isValidEmail && isValidTicket {
			// Updating the number of remaining tickets
			remainingTickets = remainingTickets - userTickets

			// Inputing users into the array
			bookings = append(bookings, firstName+" "+lastName)

			//Getting the firstNames of the users
			firstNames := []string{}
			for _, booking := range bookings {
				// separating the firstNames Fields using a whitespace
				var names = strings.Fields(booking)
				// taking the firstName of the two names at index 0
				firstNames = append(firstNames, names[0])
			}

			// Checking the remaining tickets
			if remainingTickets == 0 {
				fmt.Printf("All tickets for the %v are sold out.\n", conferenceName)
				break
			}

			fmt.Printf("The whole slice : %v.\n", bookings)
			fmt.Printf("The first slice value : %v\n", bookings[0])
			fmt.Printf("The slice length : %v.\n", len(bookings))
			fmt.Printf("slice type : %T.\n", bookings)

			fmt.Printf("Dear %v %v you have successfully booked %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v remaining tickets for %v.\n", remainingTickets, conferenceName)
			fmt.Printf("These are all the firstNames of bookings %v.\n", firstNames)
		}else{
			if !isValidUserName{
				fmt.Println("Your first and last name are less than two characters.")
			}
			if !isValidEmail{
				fmt.Println("Your email address doesn't contain @ sign.")
			}
			if !isValidTicket{
				fmt.Println("You have booked many tickets than available ones.")
			}
		}
		
	}

}
