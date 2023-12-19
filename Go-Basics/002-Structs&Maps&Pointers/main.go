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

	//one way to define a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff4343",
	}
	//anotherWay
	var colors2 map[string]string

	//one more way
	colors3 := make(map[string]string)

	colors3["white"] = "#fffff"
	delete(colors3, "white")

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {

	fmt.Printf("%v", p)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hexcode for given color:", color, "is", hex)
	}
}
