package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dkrest1/task-manager/configs"
	// "github.com/dkrest1/task-manager/routes"
	"github.com/joho/godotenv"
)

func main() {

	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env  file found!")
	}

	// init DB
	configs.InitDB()


	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in env")
	}

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}

}