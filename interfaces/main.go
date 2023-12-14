package main

import "fmt"

// Declaring structs
type englishBot struct{}
type spansishBot struct{}

// Declaring an Interface
type bot interface{
	// any struct that has function getGreeting and returns a string is 
	// promoted to be an honory member of type bot
	getGreeting() string
}

func main (){
// Initializing the bots
eb := englishBot{}
sp := spansishBot{}

// function that only accesses members of type bot
printGreeting(eb)
printGreeting(sp)

}

// function that only prints members of type bot
func printGreeting(b bot){
  fmt.Println(b.getGreeting())
}

// declaraing a function to get english greeting
func (englishBot) getGreeting() string{
    return "Hi there!"
}

// declaring a function to get spanish greeting
func (spansishBot) getGreeting() string {
	return "Hola!"
}