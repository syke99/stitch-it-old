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

	imgNm := c.Params("image")

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

	var genErr error

	if (len(colArr) != 0) && (colNum != nil) {
		if errMsg := generateExcelPattern(imgNm, colArr, colNum); errMsg != "" {
			genErrMsg := errMsg

			genErr = c.JSON(fiber.Map{"status": 500, "message": genErrMsg})
		}
	} else {
		genErr = c.JSON(fiber.Map{"status": 201, "message": "Image successfully processed and excel pattern generated"})
	}

	return genErr

}
