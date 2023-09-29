package client

import "github.com/gofiber/fiber/v2"

func Routes(rt fiber.Router) {
	rt.Post("/login", loginRoute)
	rt.Post("/guest/login", guestLoginRoute)
	rt.Post("/join", joinQueueRoute)
	rt.Post("/leave", leaveQueueRoute)
}

func loginRoute(ctx *fiber.Ctx) error {

	ctx.Status(400)
	return ctx.JSON(make(map[string]interface{}))
}

func guestLoginRoute(ctx *fiber.Ctx) error {

	ctx.Status(400)
	return ctx.JSON(make(map[string]interface{}))
}

func joinQueueRoute(ctx *fiber.Ctx) error {

	ctx.Status(400)
	return ctx.JSON(make(map[string]interface{}))
}

func leaveQueueRoute(ctx *fiber.Ctx) error {

	ctx.Status(400)
	return ctx.JSON(make(map[string]interface{}))
}
