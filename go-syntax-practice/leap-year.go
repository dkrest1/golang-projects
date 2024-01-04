package main

import "fmt"


func isLeapYear(year int) bool {
	if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
		return true
	}

	return false
}

func main() {
	var year int

	fmt.Println("Please enter year: ")
	_, err := fmt.Scanln(&year)

	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}

	if year < 0 {
		fmt.Println("Invalid input, enter a positive year")
	}

	if isLeapYear(year) {
		fmt.Println(year, "is a leap year")
	}else {
		fmt.Println(year, "is not a leap year")
	}


}