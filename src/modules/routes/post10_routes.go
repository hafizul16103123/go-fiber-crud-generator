package routes

import (
    "github.com/gofiber/fiber/v2"
    "go-crud/modules/controllers"
    "go-crud/modules/repository"
    "go-crud/modules/services"
)

func SetupPost10Routes(app *fiber.App, repo *repository.Repository) {
    service := services.NewPost10Service(repo)
    controller := controllers.NewPost10Controller(service)

    app.Get("/post10s", controller.GetAll)
    app.Get("/post10s/:id", controller.GetOne)
    app.Post("/post10s", controller.Create)
    app.Put("/post10s/:id", controller.Update)
    app.Delete("/post10s/:id", controller.Delete)
}