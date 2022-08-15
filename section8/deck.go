package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0766)
}
