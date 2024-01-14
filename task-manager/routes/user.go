package routes

import (
	"github.com/dkrest1/task-manager/controllers"
	"github.com/gorilla/mux"
)

var UserRoute *mux.Router

func userRouter( ) {

	route := mux.NewRouter()

	route.HandleFunc("/", controllers.CreateUser).Methods("POST")
	route.HandleFunc("", controllers.GetUsers).Methods("GET")
	route.HandleFunc("", controllers.FindUser).Methods("GET")
	route.HandleFunc("", controllers.UpdateUser).Methods("PATCH")
	route.HandleFunc("", controllers.DeleteUser).Methods("DELETE")

	UserRoute = route 
}