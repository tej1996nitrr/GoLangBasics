package main

import "fmt"

func main() {
	x := 15
	y := 10
	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is not less than %d", x, y)
	}
	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("different color")
	}

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("color is Blue")
	default:
		fmt.Println("different color")
	}
}
