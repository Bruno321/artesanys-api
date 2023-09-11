package service

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	UsersRoutes(app)
	ArticlesRoutes(app)
	CategoriesRoutes(app)
}

func UsersRoutes(app *fiber.App) {
	app.Get("/users", GetUsers)
	app.Get("/users/:id", GetUser)
	app.Post("/users/:id", SaveUser)
	app.Delete("/users/:id", DeleteUser)
	app.Put("/users/:id", UpdateUser)
	// app.Patch("/users/:id", service.PatchUser)
	// app.Post("/users/:id", service.ChanguePassword)
}

func ArticlesRoutes(app *fiber.App) {
	app.Get("/articles", GetArticles)
	app.Get("/articles/:id", GetArticle)
	app.Post("/articles/:id", SaveArticle)
	app.Delete("/articles/:id", DeleteArticle)
	app.Put("/articles/:id", UpdateArticle)
}

func CategoriesRoutes(app *fiber.App) {
	app.Get("/categories", GetCategories)
}
