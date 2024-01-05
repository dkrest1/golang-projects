package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	ID string 	        `json:"id,omitempty"`
	Name string	        `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Developer string     `json:"developer,omitempty"`
	Details *Details     `json:"details,omitempty"`
} 

type Details struct {
	Language string   `json:"language,omitempty"`
	Type string       `json:"type,omitempty"`
}

var app []Application

const (
	PORT = ":1909"
)

// Find one app Endpoint
func getAppEndPoint(w http.ResponseWriter, req *http.Request) {
	parms := mux.Vars(req)

	for _, item := range app {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Application{})
}

// Get all apps Endpoint
func getAppsEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(app)
}

// Create  app Endpoint
func createAppEndpoint(w http.ResponseWriter, req *http.Request) {
	var appNew Application
	_ = json.NewDecoder(req.Body).Decode(&appNew)
	app = append(app, appNew)
	json.NewEncoder(w).Encode(app)
}

// Delete 
func deleteAppEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range app {
		if item.ID == params["id"] {
			app = append(app[:index], app[index + 1:]...)
			break
		}
		
	}
	json.NewEncoder(w).Encode(app)
}

// Routing
func serverRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/app", getAppsEndPoint).Methods("GET")
	router.HandleFunc("/app/{id}", getAppEndPoint).Methods("GET")
	router.HandleFunc("/app", createAppEndpoint).Methods("POST")
	router.HandleFunc("/app/{id}", deleteAppEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(PORT, router))
}

func main() {
	app = append(app, Application{ID: "1", Name: "go-api", Description: "playing around with GO operations", Developer: "Oluwatosin", Details: &Details{Language: "Golang", Type: "Backend/API"}})
	app = append(app, Application{ID: "2", Name: "go-cli", Description: "Cli operations in GO"})
	log.Println("Sample apps populated, your api is serving at: http://127.0.0.1",PORT, "/")
	serverRequest()
}