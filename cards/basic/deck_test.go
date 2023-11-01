package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){

	d := newDeck()
	
	if len(d) != 52{
		t.Errorf("Expected deck lenght of 52, but got %v ", len(d))
	}

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected Ace of Spades at first element, but got %v instead", d[0])
	}

	if d[51] != "Rei of clubs"{

		t.Errorf("Expected Rei of clubs at last element, but got %v instead", d[51])

	}

}

func TestSaveDeckReadDeck(t *testing.T) {

	os.Remove("_decktesting.txt")

	deckToSave := newDeck()

	deckToSave.saveDeck("_decktesting.txt")

	loadedDeck, err := ReadDeck("_decktesting.txt")

	if err !=nil {
		t.Errorf("ERROR: Could not load deck: %v", err)
	}

	if len(loadedDeck) != 52{
		t.Errorf("ERROR: Could not load deck properly: %v", err)
	}

	os.Remove("_decktesting.txt")

}
