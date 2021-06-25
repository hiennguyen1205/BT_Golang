package controller

import (
	"fmt"

	"github.com/abcxyz/model"
	"github.com/abcxyz/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	listProducts := repository.Products.GetAllProducts()
	var result []*model.Product;
	for _,product := range listProducts{
		product.Rate = repository.Reviews.AverageRating(product.Id)
		result = append(result, product)
	}
	return c.JSON(result)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	productId := repository.Products.CreateProduct(product)
	return c.SendString(fmt.Sprintf("Product is created successfull with id = %d", productId))
}

func UpdateProductPatch(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repository.Products.UpdateProductPatch(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Update product successfull with id = %d", product.Id))
}

func UpSertProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	productId := repository.Products.Upsert(product)
	return c.SendString(fmt.Sprintf("Update product successfull with id = %d",productId))
}

func DeleteProductById(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repository.Products.DeleteProductById(int64(productId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("Delete product successfully")
	}
}