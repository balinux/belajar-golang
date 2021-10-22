package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke: ", i)
	}

	// continue
	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke:", x)
	}
}
