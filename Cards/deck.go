package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create new type of deck, which is slice of string

type deck []string

// Return a deck of new card
func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suits := range cardSuits {
		for _, val := range cardValue {
			cards = append(cards, val+" of "+suits)

		}
	}

	return cards
}

// Deal a hand of card
// reciever of type deck, additional parameter for handsize
// Go supports returning multiple values, so returning two deck's
func deal(d deck, handSize int) (deck, deck) {
	// Return 0:handSize, handSize:end
	return d[:handSize], d[handSize:]
}

//Convert deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",") // Join all string slices with "," to return single string
}

// Save data to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read deck from file
func newDeckFromFile(filename string) deck {
	bs, er := ioutil.ReadFile(filename) // Readfile return []byte, error -> assign them to bs and er respectively
	if er != nil {                      // if error occured er would have some value
		fmt.Println("Error: ", er)
		os.Exit(1)
	}
	str := strings.Split(string(bs), ",") // Convert []byte to string to []string
	return deck(str)
}

//Shuffle cards
func (d deck) shuffle() {
	for i := range d {

		/* Generate randomw source
		use that source to create variable of type rand, which would have random seed val
		*/
		src := rand.NewSource(time.Now().UnixNano()) // time.Now().UnixNano() -> returns random int64
		r := rand.New(src)

		pos := r.Intn(len(d) - 1)   // fixed seed value, so rand would always be same seq
		d[i], d[pos] = d[pos], d[i] //swap the elments
	}
}

// func reciever -> Any variable of type deck, gets access to function print
// d -> instance  of the variable we are working with, similar to "this" or "self"
// deck -> type, every variable of type deck can call this function on itself
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
