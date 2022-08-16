package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	cards.print()

	fmt.Print("Cards on hand:\n")
	hand.print()
	fmt.Print("Remaining cards:\n")
	remainingCards.print()
}
