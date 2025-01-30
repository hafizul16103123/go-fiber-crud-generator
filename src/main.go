package main

import (
	"context"
	"go-crud/modules/repository"
	"go-crud/modules/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	db := client.Database("go_db")

	// Initialize repository
	repo := repository.NewRepository(db)

	routes.SetupUserRoutes(app, repo)
	app.Listen(":3000")
}
