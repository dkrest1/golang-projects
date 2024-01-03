package main

import (
	"fmt"
	"math/rand"
	"time"
)

var guess int = 1

var choice int

func main()  {
	source := rand.NewSource(time.Now().UnixNano())

	randomNums := rand.New(source)

	secretNum := randomNums.Intn(10)

	fmt.Println("Guess the number")

	for {
		fmt.Println("Please enter your guess: ")
		fmt.Scanln(&choice)
		if guess > 5 {
			fmt.Println("You lose")
			break
		}else {
			if guess > secretNum {
				fmt.Println("Too big")
			}else if guess < secretNum {
				fmt.Println("Too small")
			}else {
				fmt.Println("You win!")
				break
			}
		}

		guess++
	}

}