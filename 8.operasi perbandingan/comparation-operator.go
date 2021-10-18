package main

import "fmt"

func main() {
	var num1 = 10
	var num2 = 20

	var compa bool = num1 == num2

	fmt.Println(compa)
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 != num2)
}
