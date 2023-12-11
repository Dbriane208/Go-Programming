package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T){
    // Initializing a deck of cards
	d := newDeck()

    // code to make sure we created x number of cards
	if len(d) != 16{
		t.Errorf("Expected deck length of 16, but got %v",len(d))
	}

	// code to make sure that the first card is an Ace of spades
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card of Ace of Spades, but got %v",d[0])
	}

	// code to make sure that the last card is Four of Club
	if d[len(d)-1] != "Four of Club"{
        t.Errorf("Expected last card of Four of Club, but got %v",d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
   // delete any test files from the directory 
   os.Remove("_decktesting")

   // create and save deck to the file
   deck := newDeck()
   deck.saveToFile("_decktesting")

   //load the deck from file
   loadedDeck := readNewDeckFromFile("_decktesting")

   // assert the length of the loaded deck
   if(len(loadedDeck) != 16){
	t.Errorf("Expected 16 cards in deck,got %v",len(loadedDeck))
   }

   // delete any files in current working directory with the name _decktesting
   os.Remove("_decktesting")
}