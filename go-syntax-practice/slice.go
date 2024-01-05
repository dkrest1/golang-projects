package main

import "fmt"

func main() {
	slice1 := []int{1,2,3, 4}

	slice2 := append(slice1, 5, 6)

	fmt.Println("Slice 2 comprises of slice 1 and 2: ", slice2)

	slice3 := []int{1, 2, 3}

	slice4 := make([]int, 2)

	copy(slice4, slice3)

	fmt.Println(slice3, slice4)
}