package service

import (
	"github.com/bruno321/artesanys-api/config"
	"github.com/bruno321/artesanys-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.Users
	config.DB.Find(&users)
	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.Users
	config.DB.Find(&user, id)
	return c.JSON(&user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(models.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	config.DB.Save(&user)
	return c.JSON(&user)
}

// TODO, como manejaremos los delete en artesanys
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.Users
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(500).SendString("User not found")
	}
	config.DB.Delete(&user)
	return c.SendString("User deleted")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.Users)
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(500).SendString("User not found")
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	config.DB.Save(&user)
	return c.JSON(&user)
}

// func PatchUser(c *fiber.Ctx) error {

// }

// func ChanguePassword(c *fiber.Ctx) error {

// }
