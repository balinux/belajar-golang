package main

import "fmt"

func main() {
	firstname, middlename, lastname := getFullName()

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}

func getFullName() (firstname string, middlename string, lastname string) {
	firstname = "Rio"
	middlename = "juniyantara"
	lastname = "putra"
	return
}
