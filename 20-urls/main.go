package main

import (
	"fmt"
	"net/url"
)

const myPlaceHolderURL string = "https://jsonplaceholder.typicode.com/comments?postId=1"

func main() {
	fmt.Println("Welcome to handling URLs")
	fmt.Println(myPlaceHolderURL)

	result, _ := url.Parse(myPlaceHolderURL)
	fmt.Println(result.Scheme)
	fmt.Println(result.Hostname())
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()

	fmt.Printf("The data type of qParam %T\n", qparams)
	fmt.Println(qparams["postId"])

	// Constructing the URL & was mandatory we are passing address of that URL
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "/commnets",
		RawQuery: "postId=1",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
