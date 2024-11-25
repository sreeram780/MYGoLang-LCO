package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in Go lang")

	// var ptr *int
	// fmt.Println("Value of pointer was", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of pointer was", *ptr)
	fmt.Println("Address of pointer was", ptr)

	*ptr = *ptr + 2
	fmt.Println("THe new Value was", myNumber)
}
