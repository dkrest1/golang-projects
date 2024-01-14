package routes

import (
	"net/http"

	"github.com/dkrest1/task-manager/controllers"
	"github.com/gorilla/mux"
)


func TaskRoute(taskController *controllers.TaskController ) http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/{id}", taskController.FindTask).Methods("GET")
	router.HandleFunc("/{id}", taskController.UpdateTask).Methods("PATCH")
	router.HandleFunc("/{id}", taskController.DeleteTask).Methods("DELETE")

	return router
}