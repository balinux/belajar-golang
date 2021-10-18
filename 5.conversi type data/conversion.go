package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilain64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilain64)
	fmt.Println(nilai8)

	var name = "rio"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
