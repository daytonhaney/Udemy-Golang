package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// the & operator look at this variable and
	//give me access to the memory address (ram)

	jim.updateName("Jimmy")
	jim.print()
}

// the *pointer operator take this value at the
//memory address from ^ the & and give me direct acces to the value

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
}

// the *person is a type is a description of the type
// and we are pointing for this specifically

func (p person) print() {
	fmt.Printf("%+v", p)
}
