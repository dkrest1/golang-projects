package main

import (
	"fmt"
	"strings"
)



func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	count := make(map[string]int)
	for _, word := range wordList {
		_, ok :=  count[word]

		if ok {
			count[word]++
		}else {
			count[word] = 1
		}
	}
	return count
}

func main() {
	var str = "i love golang and typescript a lot, and I will marinate my skills well on them"
	for index, value := range wordCount(str) {
		fmt.Println(index, "=>", value)
	}

	fmt.Println("The total number of characters in the  string is", len(str))
}