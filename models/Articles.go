package models

import (
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Name                  string     `json:"name" gorm:"size:50;not null" `
	Calification          float32    `json:"calification" gorm:"size:50" `
	NumberOfCalifications int        `json:"numberOfCalifications" gorm:"size:50" `
	Price                 float32    `json:"price" gorm:"size:50;not null" `
	MainImage             string     `json:"mainImage" gorm:"size:100;not null" `
	CategoriesID          int        `json:"categoriesId"`
	Categories            Categories `json:"categories" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
}

// func (article *Articles) BeforeCreate(tx *gorm.DB) (err error) {
// 	article.Calification = 0.0
// 	article.NumberOfCalifications = 0
// 	return
// }
