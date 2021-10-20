package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "rio"
	names[1] = "juniyantara"
	names[2] = "putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//
	var values = [3]int{
		90, 44, 33,
	}

	fmt.Println(values)
	fmt.Println(values[2])

	// function array
	// len = panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
}
