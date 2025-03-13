package routes

import (
	"errors"
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

func CreateOrder(c *fiber.Ctx) error {
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

func GetOrders(c *fiber.Ctx) error {
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

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("No order exists")
	}

	return nil
}

func GetOrderById(c *fiber.Ctx) error {
	id,err := c.ParamsInt("id")

	var order = models.Order{}

	if err != nil {
		return c.Status(400).JSON("Order not found")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON("Order not found")
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, "id = ?", order.UserID)
	database.Database.Db.First(&product, "id = ?", order.ProductID)

	responseOrder := CreateResponseOrder(order,CreateResponseUser(user),CreateResponseProduct(product))

	return c.Status(200).JSON(responseOrder)
}

func UpdateOrder(c *fiber.Ctx) error {
	orderId, err := c.ParamsInt("id")

	var order = models.Order{}

	if err != nil {
		return c.Status(400).JSON("Order not found")
	}

	if database.Database.Db.Find(&order, "id = ?", orderId); order.ID == 0 {
		return c.Status(400).JSON("Order not found")
	}

	var newOrder models.Order

	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := FindUser(int(newOrder.UserID), &models.User{}); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := FindProduct(int(newOrder.ProductID), &models.Product{}); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Model(&order).Updates(newOrder)

}