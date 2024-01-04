package main

import (
	"fmt"
	"os"
	"strconv"
)


func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ended before it started")
	}

	x := make([]int, stop - start)

	for i := 0; i < len(x); i++ {
		x[i] = i + 1
	}

	return x
}


func permutation(xs []int) (permuts [][]int)  {
	var rc func([]int, int)

	rc = func(a[] int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		}else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}

	rc(xs, 0)

	return permuts

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error add one numeric value ex. go run permutation.go 2")
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	noOfperms := permutation(rangeSlice(0, num))

	fmt.Println("Number of permutation", len(noOfperms))

	for i := 0; i < len(noOfperms); i++ {
		fmt.Println(noOfperms[i])
	}
}

