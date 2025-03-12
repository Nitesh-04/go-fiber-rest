package main

import (
	"log"

	"github.com/Nitesh-04/go-fiber-rest/database"
	"github.com/gofiber/fiber/v2"
)


func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome bois")
}

func main() {
	database.ConnectDb()
	
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}

