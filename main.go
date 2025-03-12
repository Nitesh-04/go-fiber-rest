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

	app.Get("api/user", routes.GetUsers)
	app.Get("api/user/:id", routes.GetUserById)
	app.Post("/api/user", routes.CreateUser)
	app.Put("/api/user/:id", routes.UpdateUser)
}

func main() {
	database.ConnectDb()
	
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

