package controllers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/services"
    "github.com/hafizul16103123/go-fiber-crud-generator/modules/models"
)

type Post11Controller struct {
    Service *services.Post11Service
}

func NewPost11Controller(service *services.Post11Service) *Post11Controller {
    return &Post11Controller{Service: service}
}

func (c *Post11Controller) Create(ctx *fiber.Ctx) error {
    var post11 models.Post11
    if err := ctx.BodyParser(&post11); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    result, err := c.Service.CreatePost11(&post11)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (c *Post11Controller) GetAll(ctx *fiber.Ctx) error {
    results, err := c.Service.GetAllPost11s()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(results)
}

func (c *Post11Controller) GetOne(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    result, err := c.Service.GetPost11ByID(id)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *Post11Controller) Update(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    var update bson.M
    if err := ctx.BodyParser(&update); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    if err := c.Service.UpdatePost11(id, update); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post11 updated successfully"})
}

func (c *Post11Controller) Delete(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    if err := c.Service.DeletePost11(id); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post11 deleted successfully"})
}