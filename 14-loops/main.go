package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sun", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	// for loop way-1
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// for loop way-2
	fmt.Println("loops in way 2")
	for i := range days {
		fmt.Println(days[i])
	}

	// for loop way-3
	fmt.Println("loops in way 3")
	for index, day := range days {
		fmt.Printf("index %v and value %v \n", index, day)
	}

	rogueValue := 1
	// for rogueValue < 10 {
	// 	fmt.Println("value is", rogueValue)
	// 	rogueValue++
	// }

	for rogueValue < 10 {
		if rogueValue == 5 {
			break
		}
		fmt.Println("value is", rogueValue)
		rogueValue++
	}
}
