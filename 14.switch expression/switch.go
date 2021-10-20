package main

import "fmt"

func main() {
	name := "rio "

	switch name {
	case "rio":
		fmt.Println("halo rio")
	case "putra":
		fmt.Println("halo putra")
	default:
		fmt.Println("halo juni")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sesuai")

	}
}
