package main

import "fmt"

func main() {
	sayHello("rio", "juniyantara")
}

func sayHello(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}
