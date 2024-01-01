package main

import "fmt"

func main( ) {
	var married = true
	var single = false
	var mother = false
	var above40 = true

	if single {
		fmt.Println("She is single and not yet married")
	}else if !single && married  {
		fmt.Println("She is married and not single")
	} else if married && !mother {
		fmt.Println("She is married and has not given birth")
	} else if married && mother {
		fmt.Println("She is married and has given birth")
	}else if single && above40 {
		fmt.Println("she is single and above 40 years ")
	}else {
		fmt.Println("We know nothing about is identity")
	}
}