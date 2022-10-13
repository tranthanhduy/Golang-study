package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#h123321",
		"blue":  "#ba12345",
	}

	// colors := make(map[int]string)

	// colors[123] = "yellow"

	// fmt.Println(colors)

	// Delete colors
	// delete(colors, 123)
	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
