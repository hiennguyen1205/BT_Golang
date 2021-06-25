package controller

import (
	"fmt"

	"github.com/abcxyz/model"
	"github.com/abcxyz/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	return c.JSON(repository.Category.GetAllCategory())
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(model.Category)

	err := c.BodyParser(&category)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	categoryId := repository.Category.CreateCategory(category)
	return c.SendString(fmt.Sprintf("Category is created successfull with id = %d", categoryId))
}

func UpdateCategoryPatch(c *fiber.Ctx) error {
	category := new(model.Category)

	err := c.BodyParser(&category)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repository.Category.UpdateCategoryPatch(category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Update category successfull with id = %d", category.Id))
}

func UpSertCategory(c *fiber.Ctx) error {
	category := new(model.Category)

	err := c.BodyParser(&category)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	categoryId := repository.Category.Upsert(category)
	return c.SendString(fmt.Sprintf("Update category successfull with id = %d",categoryId))
}

func DeleteCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repository.Category.DeleteCategoryById(int64(categoryId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("Delete category successfully")
	}
}