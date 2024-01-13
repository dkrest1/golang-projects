package routes

import (
	"github.com/dkrest1/task-manager/controllers"
	"github.com/gorilla/mux"
)

var TaskRoute *mux.Router

func taskRouter() {

	route := mux.NewRouter()

	route.HandleFunc("/", controllers.CreateTask).Methods("POST")
	route.HandleFunc("", controllers.GetTasks).Methods("GET")
	route.HandleFunc("", controllers.FindTask).Methods("GET")
	route.HandleFunc("", controllers.UpdateTask).Methods("PATCH")
	route.HandleFunc("", controllers.DeleteTask).Methods("DELETE")

	TaskRoute = route 
}