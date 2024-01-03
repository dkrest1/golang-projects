package main

import "fmt"

func main()  {
	intSlice := []int {1, 2, 3, 4, 5}

	fmt.Println(intSlice)

	firstIntSlice := intSlice[:1]

	fmt.Println(firstIntSlice)

	lastIntSlice := intSlice[len(intSlice) - 1]

	fmt.Println(lastIntSlice)

	// Remove the last element

	removeLastElement := intSlice[:len(intSlice ) - 1]

	fmt.Println(removeLastElement)

}