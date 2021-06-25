package main

import (
	"github.com/abcxyz/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func main() {
	
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// userRouter := app.Group("/api/users") //http://localhost:3000/api/users
	// routes.ConfigUserRouter(&userRouter)

	productRouter := app.Group("/api/products") //http://localhost:3000/api/products
	routes.ConfigProductRouter(&productRouter)

	categoryRouter := app.Group("/api/categories") //http://localhost:3000/api/categories
	routes.ConfigCategoryRouter(&categoryRouter)

	cartRouter := app.Group("/api/cart") //http://localhost:3000/api/cart
	routes.ConfigCartRouter(&cartRouter)

	reviewRouter := app.Group("/api/reviews") //http://localhost:3000/api/reviews
	routes.ConfigReviewRouter(&reviewRouter)

	app.Listen(":3000")
}
