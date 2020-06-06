package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday.? ")
	//use time package
	today := time.Now().Weekday()
	//Don't need to break
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
	//Display today's date and time
	fmt.Println("Today is ", time.Now())
}
