package controllers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go-crud/modules/services"
    "go-crud/modules/models"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{Service: service}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
    var user models.User
    if err := ctx.BodyParser(&user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    result, err := c.Service.CreateUser(&user)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (c *UserController) GetAll(ctx *fiber.Ctx) error {
    results, err := c.Service.GetAllUsers()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(results)
}

func (c *UserController) GetOne(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    result, err := c.Service.GetUserByID(id)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    var update bson.M
    if err := ctx.BodyParser(&update); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    if err := c.Service.UpdateUser(id, update); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    if err := c.Service.DeleteUser(id); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}