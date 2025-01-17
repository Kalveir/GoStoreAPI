package routes

import (
	"api/handlers"
	"api/middleware"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {

	//auth Route
	app.Post("/account/register", handler.RegisterAccount)
	app.Post("/login", handler.Login)

	//protected
	app.Use(middleware.Setup())

	// User Route
	app.Get("/account/profile", handler.GetProfile)
	app.Put("/account/update", handler.UpdateAccount)
	app.Delete("/account/delete", handler.DeleteAccount)
	app.Post("/logout", handler.Logout)

	// Category Route
	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)
	app.Post("/category/store", handler.CreateCategory)
	app.Put("/category/:id", handler.UpdateCategory)
	app.Delete("/category/:id", handler.DeleteCategory)

	// Product Route
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)
	app.Post("/product/store", handler.CreateProduct)
	app.Put("/product/:id", handler.UpdateProduct)
	app.Delete("/product/:id", handler.DeleteProduct)

	//Cart Route
	app.Get("/cart", handler.GetCart)
	app.Get("/cart/:id", handler.FindCart)
	app.Post("/cart/store", handler.AddCart)
	app.Put("/cart/:id", handler.UpdateCart)
	app.Delete("/cart/:id", handler.DeleteCart)

	//Order Route
	app.Get("/order", handler.GetOrder)
	app.Get("/order/:id", handler.FindOrder)
	app.Post("/order/store", handler.CreateOrder)
	app.Delete("/order/:id", handler.DeleteOrder)

	//Review Route
	app.Post("/review", handler.CreateReview)
	app.Delete("/review/:id", handler.DeleteReview)

}
