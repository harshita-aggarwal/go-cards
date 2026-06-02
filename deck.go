package main

import "fmt"

// a new type of 'deck' which 
// is a slice of string
type deck []string

// a function to create and return
// list of playing cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function to log out the contents
// of a deck of cards
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
