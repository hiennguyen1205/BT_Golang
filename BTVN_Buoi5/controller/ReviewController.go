package controller

import (
	"fmt"

	"github.com/abcxyz/model"
	"github.com/abcxyz/repository"
	"github.com/gofiber/fiber/v2"
)


func GetAllReviews(c *fiber.Ctx) error{
	return c.JSON(repository.Reviews.GetAllReviews())
}

func CreateReview(c *fiber.Ctx) error{
	review := new(model.Review)

	err := c.BodyParser(review)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	reviewId := repository.Reviews.CreateNewReview(review)
	e := AverageRating(c)
	fmt.Println(e)
	return c.SendString(fmt.Sprintf("Created review is successfull with id = %d", reviewId))
}

func UpsertReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	id := repository.Reviews.Upsert(review)
	e := AverageRating(c)
	fmt.Println(e)
	return c.SendString(fmt.Sprintf("Book with id = %d is successfully upserted", id))
}

func AverageRating(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	product, err := repository.Products.FindProductById(review.ProductId)
	if err!= nil{
		return c.SendString("Product not found")
	}
	rate := repository.Reviews.AverageRating(review.ProductId)
	product.Rate = rate
	return c.SendString("Successfull!")
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repository.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		AverageRating(c)
		return c.SendString("delete successfully")
	}
}