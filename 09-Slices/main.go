package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcomr to slices class")
	var fruits = []string{"apple", "banana", "dragon"}

	fmt.Printf("The type of fruits %T ", fruits)
	fruits = append(fruits, "Mango", "Grapes")
	fmt.Println(fruits)
	fmt.Printf("The lentgth of fruits %v", len(fruits))
	// slicing the slice
	fruits = append(fruits[1:3])
	fmt.Println(fruits)
	fruits = append(fruits[:3])
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 612
	highScores[1] = 322
	highScores[2] = 623
	highScores[3] = 435
	fmt.Println("highScores are", highScores)
	// if you uncomment below line it will gives error of index out of range
	//highScores[4] = 435

	// but if you do below thing it will build successfully
	// because append will reinitialise slice
	highScores = append(highScores, 555, 666, 777, 888, 999)
	fmt.Println("highScores are after append", highScores)

	sort.Ints(highScores)
	fmt.Println("highScores after sorting", highScores)
	fmt.Println("are highScores sorted? - Ans:", sort.IntsAreSorted(highScores))

}
