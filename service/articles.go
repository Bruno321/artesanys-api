package service

import (
	"github.com/bruno321/artesanys-api/config"
	"github.com/bruno321/artesanys-api/models"
	"github.com/gofiber/fiber/v2"
)

// TODO que directamente se us elo que viene, quitar el switch case
func GetArticles(c *fiber.Ctx) error {
	var articles []models.Articles
	filterOption := c.Query("groupBy")
	switch filterOption {
	case "cloth":
		config.DB.Where("categories_id = 1").Find(&articles)
	case "jewelry":
		config.DB.Where("categories_id = 2").Find(&articles)
	case "home":
		config.DB.Where("categories_id = 3").Find(&articles)
	case "art":
		config.DB.Where("categories_id = 4").Find(&articles)
	case "recent":
		config.DB.Order("created_at desc").Limit(3).Find(&articles)
	case "featured":
		config.DB.Order("calification desc").Limit(3).Find(&articles)
	default:
		config.DB.Find(&articles)
	}
	return c.JSON(&articles)
}

func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Articles
	config.DB.Find(&article, id)
	return c.JSON(&article)
}

func SaveArticle(c *fiber.Ctx) error {
	article := new(models.Articles)

	if err := c.BodyParser(article); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	config.DB.Save(&article)
	return c.JSON(&article)
}

// TODO, como manejaremos los delete en artesanys
func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Articles
	result := config.DB.First(&article, id)
	if result.Error != nil {
		return c.Status(500).SendString("Article not found")
	}
	config.DB.Delete(&article)
	return c.SendString("Article deleted")
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	article := new(models.Articles)
	result := config.DB.First(&article, id)
	if result.Error != nil {
		return c.Status(500).SendString("Article not found")
	}
	if err := c.BodyParser(&article); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	config.DB.Save(&article)
	return c.JSON(&article)
}
