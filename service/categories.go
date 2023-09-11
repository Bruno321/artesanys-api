package service

import (
	"github.com/bruno321/artesanys-api/config"
	"github.com/bruno321/artesanys-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Categories
	config.DB.Find(&categories)
	return c.JSON(&categories)
}
