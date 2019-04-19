package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create 'deck' which is a slice of strings
type deck []string

// function to return a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print each card in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returns hand size and remaining cards in a deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converts a deck into a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saves a deck to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// returns a deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// splits the byte into slice
	s := strings.Split(string(bs), ",")
	// convert to type of deck and return
	return deck(s)
}
