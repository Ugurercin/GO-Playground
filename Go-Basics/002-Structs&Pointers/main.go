package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	Ugur := person{firstName: "Uğur", lastName: "Erçin", contact: contactInfo{email: "ercinugur@gmail.com", zipCode: 34500}}

	Ugur.updateName("UğurPointer")
	Ugur.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {

	fmt.Printf("%v", p)
}
