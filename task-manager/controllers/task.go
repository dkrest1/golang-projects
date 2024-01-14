package controllers

import (
	"net/http"

	// "github.com/dkrest1/task-manager/models"
)


type TaskController struct{}

func NewTaskController () *TaskController {
	return &TaskController{}
}

func(c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) FindTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {

}

