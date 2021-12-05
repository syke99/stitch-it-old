package main

import (
	"bytes"
	"image"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
)

func handleFileupload(c *fiber.Ctx) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	f, err := file.Open()
	if err != nil {
		log.Println("image could not be read for processing")
		return c.JSON(fiber.Map{"status": 500, "message": "Image could not be read for proessing, bad request"})
	}

	defer f.Close()

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, f); err != nil {
		log.Println("image could not be read for processing")
		return c.JSON(fiber.Map{"status": 500, "message": "Image could not be read for proessing, bad request"})
	}

	img, _, _ := image.Decode(bytes.NewReader(buf.Bytes()))

	resize := c.FormValue("resize")

	var size string

	if resize == "false" {
		size = "default"
	} else {
		size = c.FormValue("size")
	}

	handleImageProcessing(c, img, resize, size)

	return nil
}
