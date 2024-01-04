package main

import "fmt"

func main() {
	repeatWords("I am a backend developer", 6)
}

func repeatWords(word string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(word)
	}
}