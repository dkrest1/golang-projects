package main

import "fmt"

func main() {
	isEven()
}

func isEven() {
	var num int

	fmt.Println("Please enter your number: ")
	
	_, err := fmt.Scanln(&num)

	if err != nil {
		fmt.Println("Error reading input,", err)
		return
	}

	if num % 2 == 0 {
		fmt.Println(num, "is even")
	}else {
		fmt.Println(num, "is odd")
	}
} 