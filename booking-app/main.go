package main

import (
	"fmt"
	"strconv"
)

// making the most used variables package level
// declariing variables
var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string,0)

func main() {

	greetUser()

	// %T is used to get the type of the data
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	// using an infinate for loop in Go
	for {
		
		// calling the getUserInput function
		firstName,lastName,email,userTickets := getUserInput()

		// calling the validate user inputs 
		// we're using helper.Validate... because the function is from another package
		isValidUserName,isValidEmail,isValidTicket := ValidateUserInput(firstName,lastName,email,userTickets,remainingTickets)

		// Ensuring that remaining tickets are more than userTickets
		if isValidUserName && isValidEmail && isValidTicket {

			// calling the booking ticket function
			bookTicket(firstName,lastName,email,userTickets)
			
			//Getting the firstNames of the users
			firstNames := []string{}
			firstNames = getFirstNames(firstNames)

			// Checking the remaining tickets
			if remainingTickets == 0 {
				fmt.Printf("All tickets for the %v are sold out.\n", conferenceName)
				break
			}

			// calling the conference details
            getConfDetails(firstName,lastName,email,userTickets,firstNames)

		} else {
			if !isValidUserName {
				fmt.Println("Your first and last name are less than two characters.")
			}
			if !isValidEmail {
				fmt.Println("Your email address doesn't contain @ sign.")
			}
			if !isValidTicket {
				fmt.Println("You have booked many tickets than available ones.")
			}
		}

	}

}

// creating a greeting user function
func greetUser() {
	// %v is used to get the value of the data
	fmt.Printf("Welcome to %v our booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

// function to get user input
func getUserInput() (string,string,string,uint){
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

		return firstName,lastName,email,userTickets
}


// geting the conference details
func getConfDetails(firstName string,lastName string, email string, userTickets uint, firstNames []string){
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
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

// booking ticket function
func bookTicket(firstName string,lastName string, email string,userTickets uint){
   // Updating the number of remaining tickets
   remainingTickets = remainingTickets - userTickets

   // creating a userData map
   var userData = make(map[string]string)
   userData["firstName"] = firstName
   userData["lastName"] = lastName
   userData["email"] = email
   userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)

   // Inputing users into the array
   bookings = append(bookings, userData)
   fmt.Printf("User details in a map %v.\n",userData)
}
