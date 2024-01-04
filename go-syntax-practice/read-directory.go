package main

import (
	"fmt"
	"log"
	"os"
)

func readDirectory() {
	file, err := os.Open(".")

	if err != nil {
		log.Fatalf("Failed open directory %s", err)
	}

	defer file.Close()

	fileList, err := file.ReadDir(0)

	if err != nil {
		log.Fatalf("Failed to read directory %s", err)
	}

	fmt.Printf("\nName\t\tSize\tIsDirectory  Last Modification\n")

	for _, files := range fileList {
		fmt.Printf("Name: %-15s Size: %-7v", files.Name(), files.IsDir())
	}
}

func main() {
	readDirectory()
}
