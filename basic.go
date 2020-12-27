package main

import "fmt"

func main() { // Function

	// varaible <name_of_variable> <datatype>
	//var card string = "Ace of Spades"
	//card := "Ace of String" // Infer data type
	//card := newCard()

	cards := []string{"Ace of Diamonds", newCard()} // Slice of strings [] -> declare slice, string -> data type, {} -> elements of slice
	cards = append(cards, "Five of heart")          // Append value to slice

	// for index, current element := range of elements, to slice over
	for i, card := range cards {
		fmt.Println(i, card)
	}
	//fmt.Println(cards)
}

func newCard() string { // return type comes after fun name

	return "Ace of Spades"
}
