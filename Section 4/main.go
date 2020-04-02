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
	//fmt.Printf("%+v", chat)

	chat.firstName = "Chat"
	chat.lastName = "Sumlin"

	//chat.print()
	// by default Go will be working on a copy of any data that we pass into a function
	// & is an operator in Go. &var will give the memory address of the variable referenced.
	// Rather than reference the struct directly we can point to the memory address to make sure that the correct struct is updated
	// chatPointer is now a memory address or pointer
	chatPointer := &chat
	chatPointer.updateName("John")
	chat.print()

	// You can actually just call this instead of creating a pointer like the above chatPointer := &chat
	chat.updateName("Chatfield")
	chat.print()
	// a receiver with a type of a pointer Go can coax the memory address form the struct we are operating on

	// Value Types: Use pointers to change these things in a function
	// Int, Float, String, Bool, Struct

	// Reference Type: Dont need to use pointers with these
	// slices, maps,channels, pointers, functions

	//Annotation 2020-04-02 101730.png
}

// another receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}

// *var: give me the value of whats being pointed to in Memory
// *person: a type of pointer that points to a person
// star in front of a type is different from a star in front of a var
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
