package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserId			*User
	Title 			string	`gorm:"not null"`
	Description		string
	DueDate       	time.Time
	Completed		bool

}