package controllers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go-crud/modules/services"
    "go-crud/modules/models"
)

type Post10Controller struct {
    Service *services.Post10Service
}

func NewPost10Controller(service *services.Post10Service) *Post10Controller {
    return &Post10Controller{Service: service}
}

func (c *Post10Controller) Create(ctx *fiber.Ctx) error {
    var post10 models.Post10
    if err := ctx.BodyParser(&post10); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    result, err := c.Service.CreatePost10(&post10)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (c *Post10Controller) GetAll(ctx *fiber.Ctx) error {
    results, err := c.Service.GetAllPost10s()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(results)
}

func (c *Post10Controller) GetOne(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    result, err := c.Service.GetPost10ByID(id)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *Post10Controller) Update(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    var update bson.M
    if err := ctx.BodyParser(&update); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    if err := c.Service.UpdatePost10(id, update); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post10 updated successfully"})
}

func (c *Post10Controller) Delete(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    if err := c.Service.DeletePost10(id); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post10 deleted successfully"})
}