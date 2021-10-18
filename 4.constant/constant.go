package main

import (
	"fmt"
)

func main() {
	const firstname string = "rio"
	const lastname = "juniyantara"
	const value = 1000

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

	const (
		name string = "rio"
		job         = "coding"
	)

	fmt.Println(name)
	fmt.Println(job)
}
