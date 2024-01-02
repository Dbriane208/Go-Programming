package main

import "strings"

// validating user input
func ValidateUserInput(firstName string,lastName string, email string, userTickets uint, remainingTickets uint)(bool,bool,bool){
	// validating user input
		 isValidUserName := len(firstName) >= 2 && len(lastName) >= 2
		 isValidEmail := strings.Contains(email, "@gmail.com")
		 isValidTicket := userTickets > 0 && remainingTickets >= userTickets
		 return isValidUserName,isValidEmail,isValidTicket
 }