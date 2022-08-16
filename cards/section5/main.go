package main

import "fmt"

func main() {
	animals := []string{"dog", cat(), "pig"}
	animals = append(animals, "cow")
	fmt.Println(animals)

	for i, animal := range animals {
		fmt.Println(i, animal)
	}
}

func cat() string {
	return "cat"
}
