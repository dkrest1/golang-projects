package main

import "fmt"

func main () {
	fmt.Println("Enter the three numbers: ")

	var first, second, third int

	_, err := fmt.Scanln(&first, &second, &third)

	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}

	largest := first

	if second >= largest {
		largest = second
	}

	if third >= largest {
		largest = third
	}

	fmt.Println(largest, "is the largest among the three numbers")

}