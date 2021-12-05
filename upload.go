package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func handleFileupload(c *fiber.Ctx) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", file.Filename))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	data := map[string]interface{}{

		"imageName": file.Filename,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}
