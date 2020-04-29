package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}
	for i, id := range ids {

		fmt.Println(i, id)
	}
	sum := 0
	for _, id := range ids {
		sum += id

	}
	fmt.Println("Sum", sum)
	fruits := map[string]string{"apple": "red", "orange": "orange"}
	for k, v := range fruits {
		fmt.Println(k, v)
	}
}
