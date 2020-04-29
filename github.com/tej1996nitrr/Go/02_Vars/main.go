package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var name = "Sherlock"
	var age int32 = 30
	surname := "Holmes"
	var isCool = true
	const input = "no_input"
	isCool = false
	size, email := 1.3, "sherlock.holmes@gmail.com"

	fmt.Println(surname, email)
	fmt.Println(name)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", input)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)

}
