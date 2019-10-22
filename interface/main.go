package main

import "fmt"

type frenchBot struct{}
type englishBot struct{}

func main() {

	eb := englishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	// printGreeting(fb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(fb frenchBot) {
// 	fmt.Println(fb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There !"
}

func (frenchBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Bonjour !"
}
