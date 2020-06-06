package main

//Define math package.Inside same import
import (
	"fmt"
	"math"
)

func main() {
	//Sqrt method only work with float numbers
	var num float64 = 17
	var result = math.Sqrt(num)
	fmt.Println("Original value", result)
	//Display float inside String
	fmt.Printf("Sqrt is %f .Thank you.", result)
	fmt.Println()
	//Display float in two decimal places
	fmt.Printf("Sqrt in two decimal is %.2f .Thank you.", result)
}
