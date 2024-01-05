package main

import "fmt"

func main() {
	desire := "light"

	switch desire {

		case "money":
			fmt.Println("💰💰💰💰💰💰💰")
		
		case "love":
			fmt.Println("❤️❤️❤️❤️❤️❤️❤️❤️❤️❤️❤️❤️")
		
		case "cake":
			fmt.Println("🎂🎂🎂🎂🎂🎂")
		
		case "light":
			fmt.Println("💡💡💡💡💡")
		default:
			fmt.Println("Sorry, I dont have what you desire")

	}
}