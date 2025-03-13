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
	app.Delete("/api/user/:id", routes.DeleteUser)

	app.Get("api/product", routes.GetProducts)
	app.Get("api/product/:id", routes.GetProductById)
	app.Post("/api/product", routes.CreateProduct)
	app.Put("/api/product/:id", routes.UpdateProduct)
	app.Delete("/api/product/:id", routes.DeleteProduct)

	app.Post("/api/order", routes.CreateOrder)
}

func main() {
	database.ConnectDb()
	
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

