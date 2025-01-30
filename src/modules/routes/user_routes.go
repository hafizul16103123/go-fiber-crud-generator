package routes

import (
    "github.com/gofiber/fiber/v2"
    "go-crud/modules/controllers"
    "go-crud/modules/repository"
    "go-crud/modules/services"
)

func SetupUserRoutes(app *fiber.App, repo *repository.Repository) {
    service := services.NewUserService(repo)
    controller := controllers.NewUserController(service)

    app.Get("/users", controller.GetAll)
    app.Get("/users/:id", controller.GetOne)
    app.Post("/users", controller.Create)
    app.Put("/users/:id", controller.Update)
    app.Delete("/users/:id", controller.Delete)
}