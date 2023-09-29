package server

import "github.com/gofiber/fiber/v2"

func Routes(rt fiber.Router) {
	rt.Post("/login", loginRoute)
}

func loginRoute(ctx *fiber.Ctx) error {
	return ctx.JSON(make(map[string]interface{}))
}
