package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "rio",
		"address": "bali",
	}
	// mnambah data
	person["job"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "Rio"
	book["message"] = "error"
	fmt.Println(book)
	delete(book, "message")
	fmt.Println(book)
}
