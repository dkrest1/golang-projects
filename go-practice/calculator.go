package main

import "fmt"

func main() {
	var choice int
	var firstNum float32 
	var secondNum float32

	fmt.Println("you are welcome")

	fmt.Println("Please enter the first number")
	fmt.Scanln(&firstNum)

	fmt.Println("Please enter the second number")
	fmt.Scanln(&secondNum)
	
	fmt.Println("Please input the operation number")

	fmt.Println("1 for Addition")
	fmt.Println("2 for Subtraction")
	fmt.Println("3 for Multiplication")
	fmt.Println("4 for Division")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println(firstNum, "+", secondNum, "=", firstNum + secondNum)
	}else if choice == 2 {
		fmt.Println(firstNum, "-", secondNum, "=", firstNum - secondNum)
	}else if choice == 3 {
		fmt.Println(firstNum, "*", secondNum, "=", firstNum * secondNum)
	} else if choice == 4 {
		fmt.Println(firstNum, "/", secondNum, "=", firstNum / secondNum)
	}else {
		fmt.Println("Invalid the type of operation")
	}

}