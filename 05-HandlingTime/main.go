package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	//fmt.Println(presentTime)
	// "01-02-2006 15:04:05 Monday" was standard in go lanaguage MM-dd-yyyy
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// creating date
	customDate := time.Date(2080, time.May, 16, 13, 43, 0, 0, time.UTC)
	fmt.Println(customDate)
	fmt.Println(customDate.Format("01-02-2006 15:04:05 Monday"))

}
