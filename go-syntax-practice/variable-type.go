package main

import "fmt"

func main() {

	msg := "hello"

	fmt.Printf("%v is %T\n", msg, msg)

	num := 20

	fmt.Printf("%v is %T\n", num, num)

	isboolean := 1 == 1

	fmt.Printf("Result is a %T type\n", isboolean)
}