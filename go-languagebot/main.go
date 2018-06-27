package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreenting(eb)
	printGreenting(sp)
}

func printGreenting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// A custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// A custom logic for generating an spanish greeting
	return "Hola!"
}
