package main

import "fmt"

type bot interface {
	getGreeting() string
}

type frenchBot struct{}
type englishBot struct{}

func main() {

	eb := englishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	printGreeting(fb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There !"
}

func (frenchBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Bonjour !"
}
