package handler

import (
	"api/database"
	"api/models"
	"api/service"
	"api/utils"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func GetProduct(c fiber.Ctx) error {
	product := service.GetAllProduct()
	return c.Status(200).JSON(product)
}

func FindProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := service.FindProduct(uint(id))
	return c.Status(200).JSON(product)
}

func CreateProduct(c fiber.Ctx) error {
	var product models.Product
	if err := c.Bind().Body(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validate data
	if err := utils.Validator(product); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var category models.Category
	if err := database.DB.First(&category, product.CategoryID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}
	products := service.CreateProduct(product)
	return c.Status(200).JSON(products)
}

func UpdateProduct(c fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := c.Bind().Body(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate data
	if err := utils.Validator(product); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	products := service.UpdateProduct(id, product)
	return c.Status(200).JSON(products)
}

func DeleteProduct(c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteProduct(id); err != nil {
		return c.Status(500).SendString("Failed Delete Product")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
