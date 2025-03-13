package routes

import (
	"time"

	"github.com/Nitesh-04/go-fiber-rest/database"
	"github.com/Nitesh-04/go-fiber-rest/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint    `json:"id"`
	User      User    `json:"user"`
	Product   Product `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder (c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err!=nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err:= FindUser(int(order.UserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	if err:= FindProduct(int(order.ProductID), &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}

func GetOrders (c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Find(&orders)
	
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserID)
		database.Database.Db.Find(&product, "id = ?", order.ProductID)

		responseOrder := CreateResponseOrder(order,CreateResponseUser(user),CreateResponseProduct(product))

		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}