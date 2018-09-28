package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getmarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	// person1 := Person{
	// 	firstName: "Artur",
	// 	lastName:  "Oganesyan",
	// 	city:      "Houston",
	// 	gender:    "m",
	// 	age:       25}
	//Alternative
	person1 := Person{"Artur", "Oganesyan", "Houston", "m", 25}
	person2 := Person{"Bob", "jonson", "Odessa", "f", 14}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()

	person1.getMarried("Bartosh")
	person2.getMarried("Slav")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
