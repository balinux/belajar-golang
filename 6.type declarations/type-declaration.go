package main

import "fmt"

func main() {
	type ktp string
	type married bool

	var eko ktp = "23132345252342"
	var marriedStatues married = true

	fmt.Println(eko)
	fmt.Println(marriedStatues)
}
