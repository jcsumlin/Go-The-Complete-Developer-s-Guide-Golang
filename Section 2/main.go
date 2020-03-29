package main

import "fmt"

// structs are similar to dicts from python or a plain object in JS
// Create a new definition of what a person is and the values that it can contain
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // structs can be embedded within structs
	// alternatively you can just say "contactInfo" and it will declare and define the variable within this struct
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// method of creating a new person #1
	alex := person{
		"Alex",
		"Anderson",
		contactInfo{email: "email@email.com", zipCode: 12345},
	}
	// method #2
	bob := person{firstName: "Bob", lastName: "Builder"}

	var chat person // go assignes a Zero value to all the fields in this struct.
	// Zero Values
	// string = ""
	// int, float = 0
	// bool = 0

	fmt.Println(alex)
	fmt.Println(bob)
	fmt.Println(chat)
	fmt.Printf("%+v", chat)

	chat.firstName = "Chat"
	chat.lastName = "Sumlin"

	//chat.print()

	chat.updateName("John")
	chat.print()
}

// another receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
