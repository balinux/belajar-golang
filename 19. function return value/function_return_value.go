package main

import "fmt"

func main() {
	name := getHello("rio")
	fmt.Println(name)
}

func getHello(name string) string {
	if name == "" {
		return "hello anonymous"
	} else {
		return "hello " + name
	}
}
