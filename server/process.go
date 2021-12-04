package main

import (
	"fmt"
	"image"
	"os"

	"github.com/gofiber/fiber/v2"
	dmc "github.com/syke99/go-c2dmc"
)

func handleImageProcessing(c *fiber.Ctx) error {

	logger := NewLogger()

	imageName := c.Params("imageName")

	f, err := os.Open(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		logger.log.Err(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		logger.log.Err(err)
	}

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var dmcs [][]string

	for y := 0; y < h; y++ {
		var row []string
		for x := 0; x < w; x++ {
			col := img.At(x, y)

			cb := dmc.NewColorBank()

			r, g, b := cb.RgbA(col)

			d, _ := cb.RgbToDmc(r, g, b)

			row = append(row, d)
			x++
		}
		dmcs = append(dmcs, row)
		y++
	}

	return nil
}
