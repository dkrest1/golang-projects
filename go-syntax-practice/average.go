package main

import "fmt"

func main() {
	var nums [100]int

	var temp, sum, avg int

	fmt.Println("Please enter the range of number you want to find the average: ")
	fmt.Scanln(&temp)

	for i := 0; i < temp; i++ {
		fmt.Println("Please enter the numbers to find the average: ")
		fmt.Scanln(&nums[i])

		sum += nums[i]
	}

	avg = sum / temp

	fmt.Printf("The average of the sum of %d is %d", sum, avg)
}