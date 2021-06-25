package controller

import (
	"fmt"

	"github.com/abcxyz/model"
	"github.com/abcxyz/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(repository.Users.GetAllUsers())
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	userId := repository.Users.CreateUser(user)
	return c.SendString(fmt.Sprintf("User is created successfull with id = %d", userId))
}

func UpdateUserPatch(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repository.Users.UpdateUserPatch(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Update user successfull with id = %d", user.Id))
}

func UpSertUser(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	userId := repository.Users.Upsert(user)
	return c.SendString(fmt.Sprintf("Update user successfull with id = %d",userId))
}

func DeleteUserById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repository.Users.DeleteUserById(int64(userId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("Delete user successfully")
	}
}