package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	url := "https://api.ipify.org?format=text"

	fmt.Println("Please wait while I fetch your IP address")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("This is your IP: %s\n", ip)
}