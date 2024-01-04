package main

import (
	"fmt"
	"math"
)

func main() {
	arithmetic()
	feetToMetre()
	fahrenheitToCelcius()
}

func arithmetic() {
	min := math.Min(69, 79)
	fmt.Println("The minimum between 69 and 79 is: ", min)

	max := math.Max(69, 79)
	fmt.Println("The maximum between 69 and 79 is :", max)

	fmt.Println("The square root of 16 is: ", math.Sqrt(16))

	num1 := 10
	num2 := 20
	
	fmt.Println(num1, "+", num2, "=", num1 + num2)
	fmt.Println(num1, "*", num2, "=", num1 * num2)
	fmt.Println(num1, "-", num2, "=", num1 - num2)
	fmt.Println(num1, "/", num2, "=", num1 / num2)

	u := (10 * 20) + 6

	fmt.Println("10 * 20 + 6 = ", u)
}

func feetToMetre() {
	var feetVal float32
	
	fmt.Println("Please enter feet value to convert into metre: ")
	fmt.Scanln(&feetVal)

	const meterConst = 3.28

	result := feetVal / meterConst

	fmt.Println(feetVal, "to meter = ", result)
}

func fahrenheitToCelcius() {
	var fahrenheit int

	fmt.Println("Please enter the fahrenheit temperature to convert to celcius: ")
	 
	_, err := fmt.Scanln(&fahrenheit)

	if err!= nil {
		fmt.Println("Error reading input: ", err)
	}

	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("The conversion of %v fahrenheit to celcius is %v\n", fahrenheit, celcius)
	
	
}