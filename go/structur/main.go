package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	fistName string
	lastName string
	contact  contactInfo
}

func main() {
	jim := person{
		fistName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email:   "a@a.com",
			zipCode: 2525,
		},
	}

	jim.update("abb")
	jim.print()

}

//use pointer to update name of
func (p *person) update(newFirstName string) {
	p.fistName = newFirstName
}
//use it to print person
func (p person) print() {
	fmt.Printf("%+v", p)
}
