package main

import "fmt"

func main() {
	name := "putra"

	if name == "rio" {
		fmt.Println("halo rio")
	} else if name == "putra" {
		fmt.Println("halo putra")
	} else {
		fmt.Println("halo pendatang baru")
	}

	// shord statement
	if length := len(name); length < 3 {
		fmt.Println("nama terlalu pendek")
	} else {
		fmt.Println("nama sesuai")

	}

}
