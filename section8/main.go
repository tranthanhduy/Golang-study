package main

import "fmt"

func main() {

	cards := newDeck()
	// Print all cards
	cards.print()

	// Convert from slice to string
	cards.toString()
	fmt.Println("Convert to String")
	fmt.Println(cards.toString())

	// Save above string to file
	cards.saveToFile("my_cards")

	// Read deck from file
	fmt.Println("Read deck from file")
	newCards := newDeckFromFile("my_cards")
	newCards.print()

	// Error Handling
	newCards1 := newDeckFromFile("my_cards1")
	newCards1.print()
}
