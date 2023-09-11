package models

import (
	"gorm.io/gorm"
)

type ArticlesImages struct {
	gorm.Model
	Image      string   `json:"image" gorm:"size:100;not null" `
	ArticlesID int      `json:"articlesId" gorm:"size:50;" `
	Articles   Articles `json:"articles" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
}
