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
	fmt.Printf("Name is %v and mail %v \n", sree.Name, sree.Email)
	// calling method
	sree.GetStatus()
	// updating/ manipulates prop
	sree.NewMail()
}

// struct
// make name and properties start with capital letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "udatedmail@here"
	fmt.Println("Email of this user is: ", u.Email)
}
