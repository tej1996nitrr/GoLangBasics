package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)
	fruitArray := [2]string{"Apple", "Banana"}
	fmt.Println(fruitArray)
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:1])
	

}
