package main

import "fmt"

func main(){
	// declariing variables
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	// %v is used to get the value of the data
	fmt.Printf("Welcome to %v our booking application.\n",conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	// %T is used to get the type of the data
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n",conferenceName,conferenceTickets,remainingTickets)

	// Getting user input
	var userName string
	var userTickets int

	userName = "Keen"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n",userName,userTickets)
}