package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // contactInfo contactInfo
}

func createPersonInstance() {
	john := person{"John", "Doe", contactInfo{"john.doe", 123}} // assign according to order or property
	fmt.Println(john)
	john2 := person{firstName: "John", lastName: "Doe", contactInfo: contactInfo{email: "john.doe", zipcode: 123}}
	fmt.Println(john2)
	var johnWithZeroValue person
	johnWithZeroValue.firstName = "John"
	johnWithZeroValue.print()

	johnPointer := &john
	johnPointer.updateName("Jane")
	john.print()

	// shortcut of pointer
	john.updateName("Jane2")
	john.print()
}

// pass by value
// func (p person) updateName(newName string) {
// 	p.firstName = newName
// }

// with pointer
func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
