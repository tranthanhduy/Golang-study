package main

import "fmt"

// create type "deck" which is a slice of string
type deck []string

// add
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
