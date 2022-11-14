package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // t = test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	FILENAME := "_deckTesting"
	os.Remove(FILENAME)

	deck := newDeck()
	deck.saveToFile(FILENAME)
	loadedDeck := newDeckFromFile(FILENAME)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(FILENAME)
}

func TestDealCard(t *testing.T) {

	handSize := 3
	deck := newDeck()

	hand, desk := deal(deck, handSize)

	if len(hand) > handSize {
		t.Errorf("Expected 3 cards in hand, got %v", len(hand))
	}

	if len(desk) > len(deck) - handSize {
		t.Errorf("Expected %v cards in desk, but got %v", len(deck) - 3, len(desk))
	}

}