package main

import (
	"fmt"
	"os"

	"github.com/Android-Jester/bank_queue_backend_go/src/interface/analytics"
	"github.com/Android-Jester/bank_queue_backend_go/src/interface/client"
	"github.com/Android-Jester/bank_queue_backend_go/src/interface/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var(
	mainQueue = 
)


func main() {
	app := fiber.New()

	// Middleware

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))
	// clients
	clientRoutes := app.Group("/user")
	client.Routes(clientRoutes)

	// Servers
	serverRoutes := app.Group("/server")
	server.Routes(serverRoutes)

	// Analytics
	analyticsRoutes := app.Group("/client")
	analytics.Routes(analyticsRoutes)
	if len(os.Args) < 2 {
		err := app.Listen(":3000")
		if err != nil {
			return
		}
	} else {
		port := fmt.Sprintf(":%s", os.Args[1])
		err := app.Listen(port)
		if err != nil {
			return
		}
	}
}
