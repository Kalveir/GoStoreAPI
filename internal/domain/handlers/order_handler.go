package handler

import (
	"api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// struct Request
var request struct {
	CartID []uint64 `json:"cart_id" validate:"required"`
}

/*
HANDLER GET ORDER
*/
func GetOrder(c *fiber.Ctx) error {
	// get user id
	id := c.Locals("user_id").(uint64)

	// connect to service
	cart := service.GetOrder(id)
	return c.Status(200).JSON(cart)
}

/*
HANDLER FIND ORDER
*/
func FindOrder(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64)
	id_order := c.Params("id")
	id, _ := uuid.Parse(id_order)

	// connect service
	order := service.FindOrder(id)

	// check order exist
	if order.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order Not Found",
		})
	}

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
func CreateOrder(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64) // get id user

	var totalPrice uint64 //declare variabel totalPrice
	var totalItem uint64  //declare variabel totalItem

	// Bind body request ke struct request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//get all cart
	_, cart := service.FindCart(request.CartID)

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
	totalItem = uint64(len(cart))

	//connect service
	order := service.CreateOrder(uint64(user_id), totalItem, totalPrice, cart)

	// remove cart after create order
	service.DeleteCart(request.CartID)

	return c.Status(200).JSON(order)
}

/*
HANDLER DELETE ORDER
*/
func DeleteOrder(c *fiber.Ctx) error {
	// get endpoint id & user_id
	id_order := c.Params("id")
	id, _ := uuid.Parse(id_order)
	user_id := c.Locals("user_id")

	// find Order
	order := service.FindOrder(id)

	// cek order exist
	if order.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order Not Found",
		})
	}

	// cek user order
	if order.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbiden",
		})
	}
	// connect service
	service.DeleteOrder(id)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})

}
