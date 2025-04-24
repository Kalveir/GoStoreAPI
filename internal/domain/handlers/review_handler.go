package handler

import (
	"api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

type Review = service.Review //declare type models Review

/*
HANDLER Create Review
*/
func CreateReview(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64)

	//parse body to json
	var review Review
	if err := c.BodyParser(&review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check user review
	user_review := service.FindUserReview(user_id, review.ProductID)
	if user_review.ID != 0 {
		return c.Status(403).JSON(fiber.Map{
			"message": "User Review Already Exist",
		})
	}

	//check product
	product := service.FindProduct(review.ProductID)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product Not Found",
		})
	}

	//attach user_id & product_id
	review.UserID = user_id

	//connect to service
	reviews := service.CreateReview(review)
	return c.Status(200).JSON(reviews)
}

/*
HANDLER DELETE REVIEW
*/
func DeleteReview(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user_id := c.Locals("user_id").(uint64)

	//check review
	review := service.FindReview(uint64(id))
	switch {
	case review.ID == 0:
		return c.Status(404).JSON(fiber.Map{
			"message": "Review Not Found",
		})
	case review.UserID != user_id:
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	//connect to service
	service.DeleteReview(uint64(id))
	return c.Status(200).JSON(fiber.Map{
		"message": "sucess",
	})
}
