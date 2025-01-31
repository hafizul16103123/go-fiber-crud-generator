package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/controllers"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/repository"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/services"
)

func SetupPost11Routes(app *fiber.App, repo *repository.Repository) {
    service := services.NewPost11Service(repo)
    controller := controllers.NewPost11Controller(service)

    app.Get("/post11s", controller.GetAll)
    app.Get("/post11s/:id", controller.GetOne)
    app.Post("/post11s", controller.Create)
    app.Put("/post11s/:id", controller.Update)
    app.Delete("/post11s/:id", controller.Delete)
}