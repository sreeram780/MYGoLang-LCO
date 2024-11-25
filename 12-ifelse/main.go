package main

import "fmt"

func main() {
	fmt.Println("Welcome to controlflow if else statement in go lang")

	logincount := 10

	var message string
	if logincount < 10 {
		message = "Regula User"
	} else if logincount > 10 {
		message = "hacking user"
	} else {
		message = "Exactly login in 10 times"
	}

	fmt.Println(message)

	if 9%2 == 0 {
		fmt.Println("NUmber is Even")
	} else {
		fmt.Println("NUmber is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than or equal to 10")
	}
}
