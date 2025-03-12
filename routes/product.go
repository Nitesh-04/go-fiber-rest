package routes

import (
	"errors"

	"github.com/Nitesh-04/go-fiber-rest/database"
	"github.com/Nitesh-04/go-fiber-rest/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseproduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseproduct)
}

func GetProducts (c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)

	responseproducts := []Product{}

	for _, product := range products {
		responseproduct := CreateResponseProduct(product)
		responseproducts = append(responseproducts, responseproduct)
	}

	return c.Status(200).JSON(responseproducts)

}

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("Product doesnt exist")
	}

	return nil
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product = models.Product{}

	if err != nil {
		return c.Status(400).JSON("Product not found")
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseproduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseproduct)
}

func UpdateProduct (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product = models.Product{}

	if err != nil {
		return c.Status(400).JSON("Product not found")
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type Updateproduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData Updateproduct

	if err := c.BodyParser((&updateData)); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseproduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseproduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product = models.Product{}

	if err != nil {
		return c.Status(400).JSON("Product not found")
	}

	if err:= FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err:= database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(401).JSON(err.Error())
	}

	return c.Status(200).JSON("Product deleted")
}
