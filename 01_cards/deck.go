package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of strings

type deck []string // extends all behaviour of slice of strings

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ -> we will not need to use
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// any variable of type deck can access this function
// d - copy of the deck we are working with
// something like this / self
func (d deck) print() {

	// index, current card, range of cards (iterate over every record in slice)
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// what_I_want(what_I_have)
// example: bytes := []byte("String") - converte string em array de bytes
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// perm 0666 = anyone can read or write this file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	// if nothing went wrong, it will have a value of nil
	
	if err != nil {
		// option 1 - log the error and return a call to newDeck()
		// option 2 - log the error and quit the program
		log.Fatal("Error: ", err)
	}

	s := strings.Split(string(content), ",")

	return deck(s)
}

func (d deck) shuffle() {
	// generating true random numbers
	source := rand.NewSource(time.Now().UnixNano()) // seed value
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]	
	}
}
