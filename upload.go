package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func handleFileupload(c *fiber.Ctx) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	data := map[string]interface{}{

		"image":     image,
		"imageName": file.Filename,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}
