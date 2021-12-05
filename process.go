package main

import (
	"fmt"
	"image"
	"os"

	"github.com/gofiber/fiber/v2"
)

func handleImageProcessing(c *fiber.Ctx) error {

	resize := c.FormValue("resize")

	var size string

	if resize == "false" {
		size = "default"
	} else {
		size = c.FormValue("size")
	}

	imgNm := c.Params("imageName")

	file, err := os.Open(fmt.Sprintf("./public/images/%s", imgNm))

	if err != nil {
		return c.JSON(fiber.Map{"status": 404, "message": "Image not found", "data": err})
	}

	defer file.Close()

	img, _, _ := image.Decode(file)

	var colArr [][]string

	var colNum = make(map[string]int)

	switch resize {
	case "true":
		colArr, colNum = resizeImage(img, size)
	case "false":
		colArr, colNum = noResize(img)
	}

	// TODO:
	// check to make sure colArr isn't empty
	// and then create excel pattern and delete
	// image file
	if (len(colArr) != 0) && (colNum != nil) {
		return nil
	} else {
		return c.JSON(fiber.Map{"status": 201, "message": "Image successfully processed", "data": img})
	}

}
