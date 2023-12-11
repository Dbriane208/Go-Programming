package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// creating a function to convert the slice of deck into a string
// which will have a reciever of type deck
func (d deck) toString() string{
	// using the join function to join the slice into a string using a comma
	return strings.Join([]string(d),",")
}

// creating a file to save the cards into the local repository
// we're returning an error incase something goes wrong while saving the file
func (d deck) saveToFile(filename string) error{
    return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

// creating a function to read the saved dec from the local storage
func readNewDeckFromFile(filename string) deck {
	// using read from file package in go which returns two values error and byte slice
	bs,err := ioutil.ReadFile(filename)
	if err != nil{
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Print("Error : ",err)

		// we can exit the  by passing a status code
		os.Exit(1)

	}

	// Handling the file if there's no error
	// we declare a string s which will return the slice of strings and we pass it into the deck
	s := strings.Split(string(bs),",")
	
    //return the slice of the string
	return deck(s)
}

// creating a shuffle function which will be used to generate a random number 
// this will help in efficiently randomazing the number
func (d deck) shuffle(){
	// to effectively shuffle the numbers we generate a new seed each time we start the program
	// time.Now().UnixNano() is used to generate an integer which we pass to generated a new source object
	source := rand.NewSource(time.Now().UnixNano())
	
	// we pass the new source object as the seed of our generator evertime we shuffle
	r := rand.New(source)

	// loop through the cards in the deck and then swap them
	for i := range d{
		newPosition := r.Intn(len(d)-1)
		//we swap the cards
		d[i],d[newPosition] = d[newPosition],d[i]
	}
}