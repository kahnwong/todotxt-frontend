package api

import (
	"github.com/gofiber/fiber/v2"
)

func TodoTodayController(c *fiber.Ctx) error {
	return c.JSON(getTodos("today"))
}

func TodoTinkeringController(c *fiber.Ctx) error {
	return c.JSON(getTodos("tinkering"))
}
