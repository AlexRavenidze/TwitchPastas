package controller

import "github.com/gofiber/fiber/v2"

func (h *handler) InitRouter() {
	h.fiberService.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
