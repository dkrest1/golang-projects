package main

import "fmt"

func main() {
	var sum int
	firstInput, secondInput := getUserInput()

	sum = addUserInput(firstInput, secondInput)

	fmt.Println(firstInput, "+", secondInput, "=", sum)

}

func getUserInput() (int, int) {
	var firstInput, secondInput int

	fmt.Println("Please enter first input: ")
	fmt.Scanln(&firstInput)

	fmt.Println("Please enter the second input")
	fmt.Scanln(&secondInput)

	return firstInput, secondInput

}

func addUserInput(num1, num2 int) int {
	return num1 + num2
}