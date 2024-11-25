package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch and case  in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("the random number was", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value was 1 you can open")
	case 2:
		fmt.Println("Dice value was 2 you can move 2 spots")
	case 3:
		fmt.Println("Dice value was 1 you can move 3 spots")
		fallthrough // this will execute next case as well
	case 4:
		fmt.Println("Dice value was 1 you can move 4 spots")
	case 5:
		fmt.Println("Dice value was 1 you can move 5 spots")
	case 6:
		fmt.Println("Dice value was 1 you can move 6 spots and roll dice again")
	}
}
