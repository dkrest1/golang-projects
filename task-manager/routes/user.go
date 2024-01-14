package routes

import (
	"net/http"

	"github.com/dkrest1/task-manager/controllers"
	"github.com/gorilla/mux"
)


func UserRoute(userController *controllers.UserController ) http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", userController.CreateUser).Methods("POST")
	router.HandleFunc("/", userController.GetUsers).Methods("GET")
	router.HandleFunc("/{id}", userController.FindUser).Methods("GET")
	router.HandleFunc("/{id}", userController.UpdateUser).Methods("PATCH")
	router.HandleFunc("/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}