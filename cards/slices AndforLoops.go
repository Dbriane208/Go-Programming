package main

func slicesForLoops(){
	//Initialize the function newDeck
	cards := newDeck()
	
	// print is the receiver we declared in deck.go file
	// cards.print()

	// We are recieving the cards which are of type deck and spliting them
	// with the amount we want for each handsize
	// deal(cards,5)

	// since the deal is returning two values we need to declare two variables to handle the values
	hand,remainingCards := deal(cards,5)

	// We can now print the receiver
	hand.print()
	remainingCards.print()
}
