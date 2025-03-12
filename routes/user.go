package routes

import (
	"github.com/Nitesh-04/go-fiber-rest/database"
	"github.com/Nitesh-04/go-fiber-rest/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
} // serializer

// a serializer is a mechanism for converting data structures or objects into a format that can be easily stored or transmitted, and later reconstructed

func CreateResponseUser(userModel models.User) User {
	return User {ID : userModel.ID, FirstName : userModel.FirstName, LastName : userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}