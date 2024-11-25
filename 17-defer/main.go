package main

import "fmt"

// defer will follows LIFO
// Defer in golang
// after defer in golang
// Helo defer in go
func main() {
	fmt.Println("Defer in golang")
	defer fmt.Println("Helo defer in go")
	fmt.Println("after defer in golang")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("i", i)
	}
}
