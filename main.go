package main

import (
	"log"

	"github.com/Nitesh-04/go-fiber-rest/database"
	"github.com/Nitesh-04/go-fiber-rest/routes"
	"github.com/gofiber/fiber/v2"
)


func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome bois")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	app.Post("/api/user", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

