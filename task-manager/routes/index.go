package routes

import (
	"net/http"

	"github.com/dkrest1/task-manager/controllers"
	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	router := mux.NewRouter()

	userController := controllers.NewUserController()
	taskController := controllers.NewTaskController()

	userRouter := UserRoute(userController)
	router.PathPrefix("/users").Handler(http.StripPrefix("/users", userRouter))

	taskRouter := TaskRoute(taskController)
	router.PathPrefix("/tasks").Handler(http.StripPrefix("/tasks", taskRouter))

	return router
}