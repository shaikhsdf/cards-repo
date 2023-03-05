package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a custom datatype as 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, val := range cardValues {
			// fmt.printf("%v of %v", suit, val)
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// An example of reciever function for type 'deck' to print its data print() can be any text
// also a demonstration of OO approach in go
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// demonstrates slicing
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// convert deck to type string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save deck data to a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// read the data from file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	// error handling
	if err != nil {
		fmt.Println("Error:", err)
	}

	// convert bytes data returned from readfile to deck type
	s := strings.Split(string(bs), ",")

	return deck(s)
}

// swapping and generating random numbers
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// swapping
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
