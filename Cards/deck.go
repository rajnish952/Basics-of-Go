package main

import "fmt"

// Create new type of deck, which is slice of string

type deck []string

// func reciever -> Any variable of type deck, gets access to function print
// d -> instance  of the variable we are working with, similar to "this" or "self"
// deck -> type, every variable of type deck can call this function on itself
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
