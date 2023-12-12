package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	User          	User         `gorm:"foreign_key:UserID"`
	UserID uint
	ShortUrl      	string       `gorm:"UniqueIndex;not null"`
	OriginalUrl   	string       `gorm:"not null"`
	Clicks        	uint 
	Expires       	*time.Time
}