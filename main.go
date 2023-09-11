package main

import (
	"github.com/bruno321/artesanys-api/config"
	"github.com/bruno321/artesanys-api/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.RunMigrations()

	service.Routes(app)

	app.Listen(":3000")
}
