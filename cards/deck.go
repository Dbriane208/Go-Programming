package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck [] string

// we declare a receiver
func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
}

// Returning multiple values which in this case 
// we're returning hand of deck and the remaining deck in the slice
func deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}