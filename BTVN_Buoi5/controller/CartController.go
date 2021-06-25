package controller

import (
	"fmt"

	"github.com/abcxyz/model"
	"github.com/abcxyz/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllCarts(c *fiber.Ctx) error {
	return c.JSON(repository.Carts.GetAllCarts())
}

func CreateCart(c *fiber.Ctx) error {
	cart := new(model.Cart)

	err := c.BodyParser(&cart)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	cartId := repository.Carts.CreateCart(cart)
	return c.SendString(fmt.Sprintf("Cart is created successfull with id = %d", cartId))
}

func UpdateCartPatch(c *fiber.Ctx) error {
	cart := new(model.Cart)

	err := c.BodyParser(&cart)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repository.Carts.UpdateCartPatch(cart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Update cart successfull with id = %d", cart.Id))
}

func UpSertCart(c *fiber.Ctx) error {
	cart := new(model.Cart)

	err := c.BodyParser(&cart)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	cartId := repository.Carts.Upsert(cart)
	return c.SendString(fmt.Sprintf("Update cart successfull with id = %d",cartId))
}

func DeleteCartById(c *fiber.Ctx) error {
	cartId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repository.Carts.DeleteCartById(int64(cartId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("Delete cart successfully")
	}
}