package main

import "fmt"

func main() {
	var results = 10 + 10
	fmt.Println(results)

	var (
		a = 10
		b = 20
		c = a * b
	)

	fmt.Println(c)

	// augmented assignment
	var i = 10
	i += 10
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)
}
