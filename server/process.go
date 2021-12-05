package main

import (
	"image"

	"github.com/gofiber/fiber/v2"
)

func handleImageProcessing(c *fiber.Ctx, img image.Image, resize string, size string) error {

	var colArr [][]string

	switch resize {
	case "true":
		colArr = resizeImage(img, size)
	case "false":
		colArr = noResize(img)
	}

	if len(colArr) != 0 {
		return nil
	} else {
		return c.JSON(fiber.Map{"status": 201, "message": "Image successfully processed", "data": img})
	}

}
