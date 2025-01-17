package handler

import (
	"api/service"
	"api/utils"
	"github.com/gofiber/fiber/v3"
)

// struct Request
var request struct {
	CartID []uint `json:"cart_id" validate:"required"`
}

/*
HANDLER GET ORDER
*/
func GetOrder(c fiber.Ctx) error {
	// get user id
	id := c.Locals("user_id").(uint)

	// connect to service
	cart := service.GetOrder(id)
	return c.Status(200).JSON(cart)
}

/*
HANDLER FIND ORDER
*/
func FindOrder(c fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint)
	id := c.Params("id")

	// connect service
	order := service.FindOrder(id)

	// check order exist
	// if order.ID == 0 {
	// 	return c.Status(404).JSON(fiber.Map{
	// 		"message": "Order Not Found",
	// 	})
	// }

	// check user order
	if order.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}
	return c.Status(200).JSON(order)
}

/*
HANDLER CREATE ORDER
*/
func CreateOrder(c fiber.Ctx) error {
	// get id user
	user_id := c.Locals("user_id").(uint)

	//declare variabel totalPrice
	var totalPrice uint
	//declare variabel totalItem
	var totalItem uint

	// Bind body request ke struct request
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// validate data
	if err := utils.Validator(request); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//get all cart
	cart := service.TransferCart(request.CartID)
	//check cart
	if cart == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}
	//sum total_cost
	for _, cart := range cart {

		// check cart user
		if cart.UserID != user_id {
			return c.Status(403).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
		// sum total price
		totalPrice += cart.Total_Cost

	}
	// sum total item
	totalItem = uint(len(cart))

	//connect service
	order := service.CreateOrder(uint(user_id), totalItem, totalPrice, cart)

	// remove cart after create order
	service.DeleteCart(request.CartID)

	return c.Status(200).JSON(order)
}

/*
HANDLER DELETE ORDER
*/
func DeleteOrder(c fiber.Ctx) error {
	// get endpoint id & user_id
	id := c.Params("id")
	user_id := c.Locals("user_id")

	// find Order
	order := service.FindOrder(id)

	// cek order exist
	// if order.ID == 0 {
	// 	c.Status(404).JSON(fiber.Map{
	// 		"message": "Order Not Found",
	// 	})
	// }

	// cek user order
	if order.UserID != user_id {
		c.Status(403).JSON(fiber.Map{
			"message": "Forbiden",
		})
	}
	// connect service
	service.DeleteOrder(id)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})

}
