package main

func newDeck() deck{
	//Initializing a slice of type string
	cards := deck{}

	// Create slices
	cardSuites := []string{"Spades","Diamonds","Hearts","Club"}
	cardValues := []string{"Ace","Two","Three","Four"}

	// We using and underscore to show that we declared a variable that we have not used
	for _,suit := range cardSuites{
		for _,value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}