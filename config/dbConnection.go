package config

import (
	"fmt"

	"github.com/bruno321/artesanys-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/artesanys?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

func RunMigrations() {
	db, err := InitDB()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}

	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Articles{})
	db.AutoMigrate(&models.ArticlesImages{})
}
