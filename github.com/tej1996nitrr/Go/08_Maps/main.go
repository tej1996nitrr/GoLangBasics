package main

import "fmt"

func main() {
	email := make(map[string]string)
	email["Bob"] = "bob@gmail.com"
	email["Tom"] = "tom@gmial.com"
	fmt.Println(email)
	fmt.Println(email["Bob"])
	fmt.Println(len(email))
	fruits := map[string]string{"apple": "red", "orange": "orange"}
	fmt.Println(fruits)
}
