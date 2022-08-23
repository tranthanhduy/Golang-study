package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var duy person

	duy.firstName = "Tran"
	duy.lastName = "Duy"
	duy.age = 24

	fmt.Println(duy)
	fmt.Printf("%+v", duy)
}
