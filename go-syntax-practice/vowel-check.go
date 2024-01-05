package main

import (
	"fmt"
	"strings"
)

func main() {
	vowelCheck()
}

func vowelCheck() {
	fmt.Println("Please enter your name to check if it's a vowel")

	var name string

	fmt.Scanln(&name)

	for _, char := range strings.ToLower(name) {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			fmt.Println("Your name contains a vowel character")
			return
		}
	}

	fmt.Println("Your name doesn't contain a vowel character")
}