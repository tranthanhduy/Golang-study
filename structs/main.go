package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {

	duy := person{
		firstName: "Tran",
		lastName:  "Duy",
		age:       18,
		contact: contactInfo{
			email:   "hehehihi@gmail.com",
			zipCode: 700000,
		},
	}

	fmt.Println("Name")
	duy.print()

	// fmt.Println("\nNew name 1:")
	// duyPointer := &duy
	// duyPointer.updateName("Dinivity1")
	// duy.print()

	// NOTE: Go will automatically turn normal person type to pointer person type. This is called shortcut.
	fmt.Println("\nNew name 2:")
	duy.updateName("Dinivity")
	duy.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
