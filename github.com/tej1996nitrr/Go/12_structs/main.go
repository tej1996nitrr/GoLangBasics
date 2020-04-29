package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName

}

func main() {
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// fmt.Println(person1)
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}
	fmt.Println(person2)
	// fmt.Println(person1.greet())
	person1.hasBirthday()
	// fmt.Println(person1)
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person1.greet())

}
