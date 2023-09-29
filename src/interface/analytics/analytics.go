package analytics

import "github.com/gofiber/fiber/v2"

func Routes(rt fiber.Router) {
	rt.Post("/report", reportRoute)
}

func reportRoute(ctx *fiber.Ctx) error {
	return ctx.JSON(make(map[string]interface{}))
}
