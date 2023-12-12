package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dkrest1/url-shortener/api/models"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	dbURI := os.Getenv("POSTGRES_URI")

	if dbURI == "" {
		logrus.Fatal("POSTGRES URI env not set")
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to the database")
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.User{}, &models.Url{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to auto migrate the schema")
	}


	// Start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("url shortner server is live"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("Server is running on :%s", port)

	// Listen and serve
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start the server")
	}
}
