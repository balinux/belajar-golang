package main

import "fmt"

func main() {
	var months = [...]string{
		"jan", "feb", "mar", "apr", "mey", "jun", "juli", "agus", "sept", "okt", "nov", "des",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	// slice1[0] = "mei_lain"
	// fmt.Println(months)

	// append
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "automber")
	fmt.Println(slice3)

	// make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "rio"
	newSlice[1] = "juniyantara"

	fmt.Println("\nmake slice")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("\ncopy slice")
	fmt.Println(copySlice)
}
