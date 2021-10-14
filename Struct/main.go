package main

import "fmt"

// Define type person contact structure
type contactInfo struct {
	email   string
	zipCode int
}

// Define type person structure that will exist only in our program.
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// This syntax is not good, if someone swap types above it will be issue.
	// alex := person{"Tom", "J"}

	// Second way to declare - good one
	// alex := person{firstName: "Alex", lastName: "J"}

	// Tird way to declare - good one
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "J"

	// fmt.Println(alex)

	// It will print alex struct fields.
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Me",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 80020,
		},
	}

	// & is give me the memory address of the value this variable is pointing to
	// &value - we turning value into address.
	// (*value) - we turning address into value.
	// *person - is a pointer to a type of person in memory

	// Type 1 to write it
	jim.updateName("WhoAmI")
	jim.print()

	// Type 2 to write it

}

// * - give me the value this memory address is pointing to.
// *person - this is a type description, we are wotking with a pointer to a person
// *pointerToPeson - is an operator, manipulate the value of the pointer is referencing
func (pointerToPeson *person) updateName(newFirstName string) {
	(*pointerToPeson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Reference Types - Slices, Maps, Channels, Pointers, Functions (don't worry about pointers with these)
// Value types - int, float, string, bool, structs
