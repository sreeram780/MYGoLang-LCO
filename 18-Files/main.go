package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Handlinng files in GoLang")

	content := "This was the sample file text from Sree While learning files module"
	file, err := os.Create("./myFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("the file length", length)
	// At end of file opration it should be always close
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	// file data was always in databyte format
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("the text data inside the file", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		// Panic will stops the flow and shows the error
		panic(err)
	}
}
