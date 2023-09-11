package models

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name          string `json:"name" gorm:"size:50;not null" `
	CategoryImage string `json:"categoryImage" gorm:"size:100;not null" `
	Slug          string `json:"slug" gorm:"size:50;not null" `
}
