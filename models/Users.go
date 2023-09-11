package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email      string `json:"email" gorm:"size:50;not null"`
	Password   string `gorm:"size:100;not null"`
	FirstNames string `json:"firstNames" gorm:"size:50;not null" `
	LastNames  string `json:"lastNames" gorm:"size:50;not null" `
}
