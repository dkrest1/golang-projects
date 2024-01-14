package controllers

import (

 	"net/http"

	// "github.com/dkrest1/task-manager/models"
)


type UserController struct{}

func NewUserController () *UserController {
	return &UserController{}
}

func(c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// task := models.Task{}
	
}

func(c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func(c *UserController) FindUser(w http.ResponseWriter, r *http.Request) {

}

func(c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func(c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

