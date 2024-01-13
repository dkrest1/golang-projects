package configs

import (
	"github.com/dkrest1/task-manager/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB () {
	db, err := gorm.Open(sqlite.Open("taskmanager.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Task{})

	DB = db
}
