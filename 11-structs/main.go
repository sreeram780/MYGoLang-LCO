package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in go")
	// no inheritance in golang, no parent, and child or super
	sree := User{
		Name:   "sree",
		Email:  "t@t.com",
		Status: true,
		Age:    31,
	}
	fmt.Println("Welcome to struct user:", sree)
	fmt.Printf("sree details are %+v \n", sree)
	fmt.Printf("Name is %v and mail %v", sree.Name, sree.Email)
}

// struct
// make name and properties start with capital letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
