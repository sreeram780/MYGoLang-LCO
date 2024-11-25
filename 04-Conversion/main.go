package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Println("Please reate our pizzz between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//  strconv.ParseFloat(input, 64) throws invalid syntax
	// so we will do some strings function like strconv.ParseFloat(strings.TrimSpace(input), 64)
	numrating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Thanks for rating, we will add 1 for your rating", numrating+1)
	}

}
