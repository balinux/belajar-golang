package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("data ke", counter)
		counter++
	}

	// for statement
	for i := 0; i < 10; i++ {
		fmt.Println("data ke =>", i)
	}

	// for range
	slice := []string{"aaa", "bbb", "ccc"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"budi", "marta", "sinta"}

	for index, name := range names {
		fmt.Println("index:", index, "=", name)
	}

	// looping map
	persons := make(map[string]string)

	persons["name"] = "rio"
	persons["job"] = "programmer"

	for key, persons := range persons {
		fmt.Println(key, " = ", persons)
	}
}
