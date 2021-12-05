package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func handleDeleteImage(c *fiber.Ctx, f string) error {

	patNm := strings.Split(f, ".")[0]

	if err := os.Remove("./public/images/" + f); err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": err})
	}

	if err := os.Remove("./public/patterns/" + patNm + ".xlsx"); err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": err})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
