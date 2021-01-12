package main

import "fmt"

type bot interface {
	getGreeting() string
}
type engBot struct{}
type spnBot struct{}

func main() {
	eb := engBot{}
	sb := spnBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb engBot) getGreeting() string {
	return "Hello there!"
}

func (sb spnBot) getGreeting() string {
	return "Holla!"
}
