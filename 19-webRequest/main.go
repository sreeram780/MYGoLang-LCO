package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// we can change any URL
const url = "https://youtube.com"

func main() {
	fmt.Println("Welcome to web request in golang")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The response of type %T \n", response)
	// caller's responsibility to close the connection
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Printf("The response was %v \n", content)

}
