package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greetOne()
	result := adder(3, 5)

	fmt.Println("the sum of 3, 5 is", result)

	proresult := proAdder(1, 2, 3, 4, 5)
	fmt.Println("the sum of 1, 2, 3, 4, 5 is", proresult)

	proresult, product := multireturn(1, 2, 3, 4, 5)
	fmt.Printf("the sum of 1, 2, 3, 4, 5 is %v and product %v", proresult, product)
}

// function take param and return int
func adder(a int, b int) int {
	return a + b
}

// function take array and return int
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// function take array and return int
func multireturn(values ...int) (int, int) {
	total := 0
	prod := 1
	for _, val := range values {
		total += val
		prod *= val
	}
	return total, prod
}

func greetOne() {
	fmt.Println("Welcome to evryone")
}
