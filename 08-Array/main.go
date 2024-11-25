package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays session in Golang")
	var fruitList [5]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "Dragon"

	fmt.Println(fruitList)
	fmt.Println("the length of fruitlist array", len(fruitList))

	var vegList = [3]string{"tamato", "Aloo", "Beans"}
	fmt.Println(vegList)

}
