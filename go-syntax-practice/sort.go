package main

import (
	"fmt"
	"sort"
)

func main() {
	alphabet := []string{"b", "c", "e", "y", "l", "m", "r"}
	sort.Strings(alphabet)

	fmt.Println("Sorted Alphabet", alphabet)

	integers := []int{45, 32, 56, 67, 89}
	sort.Ints(integers)

	fmt.Println(integers)

	a := sort.IntsAreSorted(integers)

	fmt.Println("Integers are sorted?", a)

	b := sort.StringsAreSorted(alphabet)

	fmt.Println("Alphabet are sorted?", b)
	
}