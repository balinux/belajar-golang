package main

import "fmt"

func main() {
	firstname, _, lastname := getFullname()
	fmt.Println(firstname, lastname)
}

func getFullname() (string, string, string) {
	return "rio", "juniyantara", "putra"
}
