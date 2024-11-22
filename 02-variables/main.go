package main

import "fmt"

// constant of global variable
const apiToken string = ""

// constant of package variable
const ApiToken string = ""

func main() {
	var userName string = "Sree"
	fmt.Println(userName)
	fmt.Printf("Variable is type of %T \n", userName)

	var isUserLoggedIn bool = false
	fmt.Println(isUserLoggedIn)
	fmt.Printf("Variable is type of %T \n", isUserLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of %T \n", smallVal)

	var smallFloat float32 = 4.8
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of %T \n", smallFloat)

	var intval int
	fmt.Println(intval)
	fmt.Printf("Variable is type of %T \n", intval)

	var strVal string
	fmt.Println(strVal)
	fmt.Printf("Variable is type of %T \n", strVal)

	// this will works only with in the method or fuction
	impicitName := "This was implicit name"
	fmt.Println(impicitName)
	fmt.Printf("Variable is type of %T \n", impicitName)

}
