package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReviews(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func GetReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	review, err := repo.Reviews.FindReviewById(int64(id))
	// fmt.Println(review)
	// fmt.Println(err)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(review)
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		e := AverageRating(c)
		fmt.Println(e)
		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	fmt.Println(err)
	// if error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	if _, err := repo.Books.FindBookById(review.BookId); err != nil {
		return c.JSON(fiber.Map{
			"message": "That book does not exist",
		})
	}

	reviewId := repo.Reviews.CreateNewReview(review)
	return c.SendString(fmt.Sprintf("New book is created successfully with id = %d", reviewId))

}

func UpdateReview(c *fiber.Ctx) error {
	updatedReview := new(model.Review)

	err := c.BodyParser(&updatedReview)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Reviews.UpdateReview(updatedReview)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Book with id = %d is successfully updated", updatedReview.Id))

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

	id := repo.Reviews.Upsert(review)
	e := AverageRating(c)
	fmt.Println(e)
	return c.SendString(fmt.Sprintf("Book with id = %d is successfully upserted", id))
}

func AverageRating(c *fiber.Ctx) error {
	listAllReview := repo.Reviews.GetAllReviews()
	for _, detailReview := range listAllReview {
		book, err := repo.Books.FindBookById(int64(detailReview.BookId))
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Not found book for this id",
			})
		}
		result := repo.Reviews.AverageRating()
		book.Rating = result[int64(detailReview.BookId)]
		repo.Books.UpdateBook(book)
	}
	return c.SendString("Successfull!")
}
