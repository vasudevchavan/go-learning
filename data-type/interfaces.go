package main

import "fmt"

type botEngish struct{}
type botSpanish struct{}

// Define a common interface for both bot types
type pBot interface {
	getGreeting() string
}

func (eb botEngish) getGreeting() string {
	return "Hello, friend!"
}

func (sb botSpanish) getGreeting() string {
	return "Hola, amigo!"
}

func (eb botEngish) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb botSpanish) printGreeting() {
	fmt.Println(sb.getGreeting())
}

// Function that accepts the common interface
func cPrintGreeting(b pBot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := botEngish{}
	sb := botSpanish{}

	eb.printGreeting()
	sb.printGreeting()

	// Using the common interface
	cPrintGreeting(eb)
	cPrintGreeting(sb)

}
