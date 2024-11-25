package main

import "fmt"

func main() {
	fmt.Println("welcome to maps demo in Golang")
	// make(map[keyType]valueType)
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println("js stand for", languages["js"])

	// deleting key from maps
	delete(languages, "py")
	fmt.Println(languages)

	// loops are intresting in golang

	for index, value := range languages {
		fmt.Println(index)
		fmt.Println(value)
	}

	for _, value := range languages {
		fmt.Println(value)
	}

	for index, _ := range languages {
		fmt.Println(index)
	}
}
